package recent

import (
	"embed"
	"html/template"
	"slices"
	"strings"
//	"fmt"
	_ "embed"

	. "github.com/emad-elsaid/xlog"
)

//go:embed templates
var templates embed.FS

func init() {
	RegisterExtension(Recent{})
}

type Recent struct{}

func (Recent) Name() string { return "recent" }
func (Recent) Init() {
	Get(`/+/recent`, recentHandler)
	Get(`/+/recent/created`, ctimeHandler)
	RegisterBuildPage("/+/recent", true)
	RegisterBuildPage("/+/recent/created", true)
	RegisterTemplate(templates, "templates")
	RegisterLink(func(Page) []Command { return []Command{links{}} })
}

func recentHandler(r Request) Output {
	rp := slices.Clone(Pages(r.Context()))

	rp = slices.DeleteFunc(rp, func(a Page) bool {
		return IsIgnoredPath(a.Name())
	})

	slices.SortFunc(rp, func(a, b Page) int {
		if modtime := b.ModTime().Compare(a.ModTime()); modtime != 0 {
			return modtime
		}

		return strings.Compare(a.Name(), b.Name())
	})

	return Render("recent", Locals{
		"page":  DynamicPage{NameVal: "Recent"},
		"pages": rp,
	})
}

func ctimeHandler(r Request) Output {
	rp := slices.Clone(Pages(r.Context()))

	rp = slices.DeleteFunc(rp, func(a Page) bool {
		return IsIgnoredPath(a.Name())
	})

	slices.SortFunc(rp, func(a, b Page) int {
		a_yr := strings.HasPrefix(a.Name(), "20");
		b_yr := strings.HasPrefix(b.Name(), "20");
		if (a_yr && b_yr) {
			return strings.Compare(b.Name(), a.Name())
		}
		if (!a_yr && !b_yr) {
			return strings.Compare(a.Name(), b.Name())
		}
		if a_yr {
			return -1
		}
		return 1
	})

	return Render("recent", Locals{
		"page":  DynamicPage{NameVal: "Recent"},
		"pages": rp,
		"ctime": true,
	})
}


type links struct{}

func (l links) Icon() string { return "" }
func (l links) Name() string { return "Recent" }
func (l links) Attrs() map[template.HTMLAttr]any {
	return map[template.HTMLAttr]any{
		"href": "/+/recent",
		"accesskey": "r",
	}
}
func (l links) Label() map[string]string {
		return map[string]string {
			"labelStart": "",
			"labelAccel": "R",
			"labelEnd": "ecent",
	}
}
