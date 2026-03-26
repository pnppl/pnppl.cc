package newest

import (
	"fmt"
	"html/template"
	"strings"
/*	"io/fs"
	"path"
	"path/filepath"
	"slices"
*/
	"github.com/emad-elsaid/xlog"
	"github.com/emad-elsaid/xlog/extensions/shortcode"
)

func init() {
	xlog.RegisterExtension(Newest{})
}

type Newest struct{}

func (Newest) Name() string { return "newest" }
func (Newest) Init() {
	shortcode.RegisterShortCode("newest", shortcode.ShortCode{Render: newestShortcode})
}

func newestShortcode(in xlog.Markdown) template.HTML {
	newestName := xlog.Newest()
	if newestName == "" {
		return template.HTML(fmt.Sprintf("[placeholder. currently only works when built, not when tested live, due to assumptions about build order and available context]"))
	}
	p := xlog.NewPage(newestName)
	props := xlog.Properties(p)
	title := newestName
	if v, ok := props["title"]; ok && v != nil {
		if s, ok := v.Value().(string); ok {
			title = s
		}
	}
	var content []string
	i := 0
	for line := range strings.Lines(string(p.Render())) {
		if i == 5 { break }
		content = append(content, line)
		i++
	}
	contents := strings.TrimSpace(strings.Join(content[0:], ""))
	return template.HTML(fmt.Sprintf(`<h2 class="excerpt">Newest post: <a href="%s" class="excerpt">%s</a></h2> <blockquote class="excerpt">%s<a href="%s" class="excerpt-more">...</a></blockquote>`, newestName, title, contents, newestName))
}
