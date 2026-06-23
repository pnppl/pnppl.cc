package autolink_pages

import (
	"fmt"
	"slices"
	"strings"
	"regexp"
	"github.com/emad-elsaid/xlog"
	"github.com/emad-elsaid/xlog/markdown/ast"
	"github.com/emad-elsaid/xlog/markdown/renderer"
	"github.com/emad-elsaid/xlog/markdown/util"
)

type pageLinkRenderer struct{}

func (h *pageLinkRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(KindPageLink, render)
}

func render(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering {
		n := node.(*PageLink)
		url := n.page.Name()

		// identify linked document by its hashtags and put definitions in tooltip
		class := "autolink"
		tooltip := ""
		parent := xlog.NewPage(url)
		props := xlog.Properties(parent)
		if v, ok := props["tags"]; ok && v != nil {
			if slices.Contains(v.Value().([]string), "glossary") {
				class = class + " glossary"
				lines := strings.Split(string(parent.Content()), "\n")
				for _, line := range lines {
					if len(line) > 0 {
						// bail out when we hit a hashtag or footnote (or heading)
						if strings.HasPrefix(line, "#") || strings.HasPrefix(line, `[^`) {
							break
						}
						tooltip = tooltip + string(line) + " "
					}
				}
				// delete [^...]
				footnoterefRE := regexp.MustCompile(`\[\^([^]]+)\]`)
				tooltip = footnoterefRE.ReplaceAllString(tooltip, "")
				// [[a|...]] -> a
				linkRE := regexp.MustCompile(`\[\[([^|]+)\|[^]]+\]\]`)
				tooltip = linkRE.ReplaceAllString(tooltip, "$1")
			}
		}
		if len(tooltip) == 0 {
			tooltip = "mention"
		}
		fmt.Fprintf(w,
			`<a href="/%s" class="%s" title="%s">`,
			util.EscapeHTML(util.URLEscape([]byte([]byte(url)), false)),
			class,
			util.EscapeHTML([]byte(strings.Trim(tooltip, " \t"))),
//			tooltip,
		)

		if total, done := countTodos(n.page); total > 0 {
			isDone := ""
			if total == done {
				isDone = "is-success"
			}
			fmt.Fprintf(w, `<span class="tag is-rounded %s">%d/%d</span> `, isDone, done, total)
		}
	} else {
		w.WriteString(`</a>`)
	}

	return ast.WalkContinue, nil
}
