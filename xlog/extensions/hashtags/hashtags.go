package hashtags

import (
	"context"
	"embed"
	"fmt"
	"html/template"
	"slices"
	"strings"
	"sync"
	"unicode"
	"unicode/utf8"
	"unique"
	_ "log/slog"

	. "github.com/emad-elsaid/xlog"
	"github.com/emad-elsaid/xlog/extensions/shortcode"
	"github.com/emad-elsaid/xlog/markdown/ast"
	"github.com/emad-elsaid/xlog/markdown/parser"
	"github.com/emad-elsaid/xlog/markdown/renderer"
	"github.com/emad-elsaid/xlog/markdown/text"
	"github.com/emad-elsaid/xlog/markdown/util"
)

//go:embed templates
var templates embed.FS

var h Hashtags

func init() {
	h = Hashtags{
		pages: make(map[Page][]*Hashtag),
	}

	RegisterExtension(&h)
}

type Hashtags struct {
	pages map[Page][]*Hashtag
	mu    sync.Mutex
}

func (*Hashtags) Name() string { return "hashtags" }
func (h *Hashtags) Init() {
	Get(`/+/tags`, h.tagsHandler)
	Get(`/+/tag/{tag}`, h.tagHandler)
	RegisterWidget(WidgetAfterView, 1, h.relatedPages)
	RegisterBuildPage("/+/tags", true)
	RegisterLink(links)
	RegisterTemplate(templates, "templates")
	shortcode.RegisterShortCode("hashtag-pages", shortcode.ShortCode{Render: h.hashtagPages})
	shortcode.RegisterShortCode("hashtag-pages-grid", shortcode.ShortCode{Render: h.hashtagPagesGrid})

	Listen(PageChanged, h.PageChanged)
	Listen(PageDeleted, h.PageDeleted)

	MarkdownConverter().Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(&Hashtag{}, 0),
	))
	MarkdownConverter().Parser().AddOptions(parser.WithInlineParsers(
		util.Prioritized(&Hashtag{}, 999),
	))
	GemtextConverter().Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(&HashtagGemtext{}, 0),
	))

	RegisterProperty(properties)
}

func (h *Hashtags) PageChanged(p Page) error {
	delete(h.pages, p)

	return nil
}

func (h *Hashtags) PageDeleted(p Page) error {
	return h.PageChanged(p)
}

func (h *Hashtags) hashtagsFor(p Page) []*Hashtag {
	h.mu.Lock()
	defer h.mu.Unlock()

	if tags, ok := h.pages[p]; ok {
		return tags
	}

	_, tree := p.AST()
	tags := FindAllInAST[*Hashtag](tree)
	h.pages[p] = tags

	return tags
}

func (h *Hashtags) tagsHandler(r Request) Output {
	tags := map[string][]Page{}
	var lck sync.Mutex

	EachPage(r.Context(), func(a Page) {
		set := map[string]bool{}
		_, tree := a.AST()
		hashes := FindAllInAST[*Hashtag](tree)
		for _, v := range hashes {
			val := strings.ToLower(string(v.value))

			// don't use same tag twice for same page
			if _, ok := set[val]; ok {
				continue
			}

			set[val] = true

			lck.Lock()
			if ps, ok := tags[val]; ok {
				tags[val] = append(ps, a)
			} else {
				tags[val] = []Page{a}
			}
			lck.Unlock()
		}
	})
	for tag, _ := range tags {
		tags[tag] = PageTitleSort(tags[tag])
	}

	return Render("tags", Locals{
		"page": DynamicPage{NameVal: "Hashtags"},
		"tags": tags,
	})
}

func (h *Hashtags) tagHandler(r Request) Output {
	tag := r.PathValue("tag")
	pages := h.tagPages(r.Context(), tag)
	pages = PageTitleSort(pages)
	return Render("tag", Locals{
		"page":  DynamicPage{NameVal: "#" + tag},
		"pages": pages,
	})
}

func (h *Hashtags) tagPages(ctx context.Context, hashtag string) []Page {
	uniqHandle := unique.Make(strings.ToLower(hashtag))

	return MapPage(ctx, func(p Page) Page {
		if p.Name() == Config.Index {
			return nil
		}

		tags := h.hashtagsFor(p)
		for _, t := range tags {
			if uniqHandle == t.unique {
				return p
			}
		}

		return nil
	})
}

func (h *Hashtags) relatedPages(p Page) template.HTML {
	if p.Name() == Config.Index {
		return ""
	}

	_, tree := p.AST()
	found_hashtags := FindAllInAST[*Hashtag](tree)
	hashtags := map[unique.Handle[string]]bool{}
	for _, v := range found_hashtags {
		hashtags[v.unique] = true
	}

	// this stuff runs concurrently so we need this to prevent a crash from multiple simultaneous writes to pages_with_counts. there's probably a better way to do it
	var mu sync.Mutex
	pages_with_counts := map[Page]int{}
	pages := MapPage(context.Background(), func(rp Page) Page {
		if rp.Name() == p.Name() {
			return nil
		}

		_, tree := rp.AST()
		page_hashtags := FindAllInAST[*Hashtag](tree)
		page_shared_count := 0
		for _, h := range page_hashtags {
			if _, ok := hashtags[h.unique]; ok {
				page_shared_count++
			}
		}
		if page_shared_count > 0 {
			mu.Lock()
			pages_with_counts[rp] = page_shared_count
			mu.Unlock()
			return rp
		}

		return nil
	})

	// sort by tags in common
	slices.SortFunc(pages, func(a, b Page) int {
		if pages_with_counts[a] == pages_with_counts[b] {
			return strings.Compare(strings.ToLower(GetTitle(a)), strings.ToLower(GetTitle(b)))
		}
		if pages_with_counts[a] > pages_with_counts[b] {
			return -1
		}
		return 1
	})

	// shrink
	limit := 5
	max_shared := pages_with_counts[pages[0]]
	min_shared := pages_with_counts[pages[len(pages)-1]]
	// remove most irrelevant pages
	// check max != min because we don't want to delete the whole list (eg we still want relateds for single-tag posts)
	if len(pages) > limit && max_shared != min_shared {
		pages = slices.DeleteFunc(pages, func(a Page) bool {
			return pages_with_counts[a] == min_shared
		})
	}
	if len(pages) > limit {
		pages = pages[:limit]
	}

	return Partial("related-hashtags-pages", Locals{
		"pages": pages,
	})
}

func (h *Hashtags) hashtagPages(hashtag Markdown) template.HTML {
	hashtag_value := strings.Trim(string(hashtag), "# \n")
	pages := h.tagPages(context.Background(), hashtag_value)

	slices.SortFunc(pages, func(a, b Page) int {
		if modtime := b.ModTime().Compare(a.ModTime()); modtime != 0 {
			return modtime
		}

		return strings.Compare(a.Name(), b.Name())
	})

	output := Partial("hashtag-pages", Locals{"pages": pages})
	return template.HTML(output)
}

func (h *Hashtags) hashtagPagesGrid(hashtag Markdown) template.HTML {
	hashtag_value := strings.Trim(string(hashtag), "# \n")
	pages := h.tagPages(context.Background(), hashtag_value)

	slices.SortFunc(pages, func(a, b Page) int {
		if modtime := b.ModTime().Compare(a.ModTime()); modtime != 0 {
			return modtime
		}

		return strings.Compare(a.Name(), b.Name())
	})

	output := Partial("hashtag-pages-grid", Locals{"pages": pages})
	return template.HTML(output)
}

func links(Page) []Command {
	return []Command{link{}}
}

type link struct{}

func (l link) Icon() string { return "" }
func (l link) Name() string { return "Hashtags" }
func (l link) Attrs() map[template.HTMLAttr]any {
	return map[template.HTMLAttr]any{
		"href": "/+/tags",
		"accesskey": "t",
	}
}
func (l link) Label() map[string]string {
		return map[string]string {
			"labelStart": "Hash",
			"labelAccel": "t",
			"labelEnd": "ags",
	}
}

type Hashtag struct {
	ast.BaseInline
	value  []byte
	unique unique.Handle[string]
}

func (h *Hashtag) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(KindHashtag, renderHashtag)
}

func (h *Hashtag) Trigger() []byte {
	return []byte{'#'}
}

func (h *Hashtag) Parse(parent ast.Node, block text.Reader, pc parser.Context) ast.Node {
	l, _ := block.PeekLine()
	if len(l) < 1 {
		return nil
	}

	var line string = string(l)

	var i int
	for ui, c := range line {
		if ui == 0 {
			i += utf8.RuneLen(c)
			continue
		}

		if !(unicode.In(c, unicode.Letter, unicode.Number, unicode.Dash) || c == '_') || unicode.IsSpace(c) {
			break
		}

		i += utf8.RuneLen(c)
	}
	if i > len(line) || i == 1 {
		return nil
	}

// no hashtags that are all digits
	for ui, c := range line[1:i] {
		if c < '0' || c > '9' {
			break
		}
		if ui == len(line[1:i]) - 1 {
			return nil
		}
	}

	block.Advance(i)
	tag := line[1:i]
	return &Hashtag{
		value:  []byte(tag),
		unique: unique.Make(strings.ToLower(tag)),
	}
}

func (h *Hashtag) Dump(source []byte, level int) {
	m := map[string]string{
		"value": fmt.Sprintf("%#v", h.value),
	}
	ast.DumpHelper(h, source, level, m, nil)
}

var KindHashtag = ast.NewNodeKind("Hashtag")

func (h *Hashtag) Kind() ast.NodeKind {
	return KindHashtag
}

type HashtagGemtext struct {
	*Hashtag
}

func (h *HashtagGemtext) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(KindHashtag, renderHashtagGemtext)
}

// HTML
func renderHashtag(writer util.BufWriter, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
	return renderHashtagGeneric(writer, source, n, entering, "html")
}

func renderHashtagGemtext(writer util.BufWriter, source []byte, n ast.Node, entering bool) (ast.WalkStatus, error) {
	return renderHashtagGeneric(writer, source, n, entering, "gemtext")
}

func renderHashtagGeneric(writer util.BufWriter, source []byte, n ast.Node, entering bool, format string) (ast.WalkStatus, error) {
	if !entering || n.Kind() != KindHashtag {
		return ast.WalkContinue, nil
	}

	tag := n.(*Hashtag)
	switch format {
		case "html":
			fmt.Fprintf(writer, `<a href="/+/tag/%s" class="tag" id="tag-%s" name="tag-%s" data-pagefind-meta="tag-%s:%s" data-pagefind-filter="hashtag:%s">#%s</a>`, tag.value, tag.value, tag.value, tag.value, tag.value, tag.value, tag.value)
		case "gemtext":
			fmt.Fprintf(writer, `%s=> /+/tag/%s #%s`, "\n", tag.value, tag.value)
	}
	RegisterBuildPage(fmt.Sprintf("/+/tag/%s", tag.value), true)
	RegisterBuildPage(fmt.Sprintf("/+/tag/%s", strings.ToLower(string(tag.value))), true)
	return ast.WalkContinue, nil
}
