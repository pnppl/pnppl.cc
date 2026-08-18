package all

import (
	"embed"
	"html/template"
	"slices"
	"strings"
	"regexp"
	_ "embed"

	. "github.com/emad-elsaid/xlog"
)

//go:embed templates
var templates embed.FS

func init() {
	RegisterExtension(All{})
}

type All struct{}

func (All) Name() string { return "all" }
func (All) Init() {
	Get(`/+/all`, allHandler)
	Get(`/+/all/filename`, filesort)
	RegisterBuildPage("/+/all", true)
	RegisterBuildPage("/+/all/filename", true)
	RegisterTemplate(templates, "templates")
	RegisterLink(func(Page) []Command { return []Command{links{}} })
}

func allHandler(r Request) Output {
	pages := slices.Clone(Pages(r.Context()))

	pages = PageTitleSort(pages)

	return Render("all", Locals{
		"page":  DynamicPage{NameVal: "All"},
		"pages": pages,
	})
}

func filesort(r Request) Output {
	rp := slices.Clone(Pages(r.Context()))

	rp = slices.DeleteFunc(rp, func(a Page) bool {
		return IsIgnoredPath(a.Name())
	})

	re := regexp.MustCompile(`^[A-z]`)
	slices.SortFunc(rp, func(a, b Page) int {
		a_yr := ! re.Match([]byte(a.Name()))
		b_yr := ! re.Match([]byte(b.Name()))
		if (a_yr && b_yr) || (!a_yr && !b_yr) {
			return strings.Compare(a.Name(), b.Name())
		}
		if a_yr {
			return 1
		}
		return -1
	})

	return Render("all", Locals{
		"page":  DynamicPage{NameVal: "All"},
		"pages": rp,
		"filesort": true,
	})
}



type links struct{}

func (l links) Icon() string { return "" }
func (l links) Name() string { return "All" }
func (l links) Attrs() map[template.HTMLAttr]any {
	return map[template.HTMLAttr]any{
		"href": "/+/all",
		"accesskey": "a",
	}
}
func (l links) Label() map[string]string {
		return map[string]string {
			"labelStart": "",
			"labelAccel": "A",
			"labelEnd": "ll",
	}
}
