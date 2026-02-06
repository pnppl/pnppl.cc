package search

import (

	"embed"
	"html/template"

	_ "embed"

	. "github.com/emad-elsaid/xlog"
)

//go:embed templates
var templates embed.FS

func init() {
	RegisterExtension(Search{})
}

type Search struct{}

func (Search) Name() string { return "search" }
func (Search) Init() {
	Get(`/+/search`, searchHandler)
	RegisterBuildPage("/+/search", true)
	RegisterTemplate(templates, "templates")
	RegisterLink(func(Page) []Command { return []Command{links{}} })
}

type links struct{}
func (l links) Icon() string { return "" }
func (l links) Name() string { return "Search" }
func (l links) Attrs() map[template.HTMLAttr]any {
	return map[template.HTMLAttr]any{
		"href": "/+/search",
	}
}

func searchHandler(r Request) Output {
	return Render("search", Locals{
		"page":  DynamicPage{NameVal: "Search"},
		"pages": nil,
	})
}
