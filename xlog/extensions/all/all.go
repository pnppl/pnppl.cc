package all

import (
	"embed"
	"html/template"
	"slices"
	"strings"

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
	RegisterBuildPage("/+/all", true)
	RegisterTemplate(templates, "templates")
	RegisterLink(func(Page) []Command { return []Command{links{}} })
}

func allHandler(r Request) Output {
	rp := Pages(r.Context())
	slices.SortFunc(rp, func(a, b Page) int {
//		if modtime := b.ModTime().Compare(a.ModTime()); modtime != 0 {
//			return modtime
//		}
		nameA := a.Name()
		nameB := b.Name()

		content := a.Content()
		lines := strings.Split(string(content), "\n")
		firstLine := strings.TrimSpace(lines[0])
		normalizedNameA := strings.Replace(strings.Replace(firstLine, "# ", "", 1), " #", "", 1)

		content = b.Content()
		lines = strings.Split(string(content), "\n")
		firstLine = strings.TrimSpace(lines[0])
		normalizedNameB := strings.Replace(strings.Replace(firstLine, "# ", "", 1), " #", "", 1)

		if len([]byte(normalizedNameA)) > 0 {
			nameA = normalizedNameA
		}
		if len([]byte(normalizedNameB)) > 0 {
			nameB = normalizedNameB
		}
		return strings.Compare(strings.ToLower(nameA), strings.ToLower(nameB))
//		return strings.Compare(a.Name(), b.Name())
	})

	return Render("all", Locals{
		"page":  DynamicPage{NameVal: "All"},
		"pages": rp,
	})
}

type links struct{}

func (l links) Icon() string { return "fa-solid fa-clock-rotate-left" }
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
