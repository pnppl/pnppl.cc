package shortcode

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"
	. "github.com/emad-elsaid/xlog"
)

type ShortCode struct {
	Render  func(Markdown) template.HTML
	Default string
}

func render(i Markdown) string {
	var b bytes.Buffer
	MarkdownConverter().Convert([]byte(i), &b)
	return b.String()
}

func container(cls string, content Markdown) template.HTML {
	tpl := `<article class="message %s"><div class="message-body">%s</div></article>`
	return template.HTML(fmt.Sprintf(tpl, cls, render(content)))
}

var shortcodes = map[string]ShortCode{
//	"info":    {Render: func(c Markdown) template.HTML { return container("is-info", c) }},
//	"success": {Render: func(c Markdown) template.HTML { return container("is-success", c) }},
//	"warning": {Render: func(c Markdown) template.HTML { return container("is-warning", c) }},
	"!":   {Render: func(c Markdown) template.HTML {
		text := render(c)
		text = strings.TrimPrefix(text, "<p>")
		text = strings.TrimSuffix(text, "</p>\n")
		excl := `<span class="is-hidden">!!<br><br></span>`
		return template.HTML(fmt.Sprintf(`<div class="message is-danger"><div class="message-body"><b>%s%s</b></div></div>`, excl, text)) }},
//	"!!":   {Render: func(c Markdown) template.HTML { return container("is-danger", c) }},
/* This would be so convenient but it breaks the footnotes and I CBA to figure it out rn
	"deets":   {Render: func(c Markdown) template.HTML {
		text := render(c)
		text = strings.TrimPrefix(text, "<p>")
		text = strings.TrimSuffix(text, "</p>\n")
		summary, _, _ := strings.Cut(text, "\n")
		return template.HTML(fmt.Sprintf(`<details><summary>%s</summary>%s</details>`, summary, text)) }},
*/
}

func RegisterShortCode(name string, shortcode ShortCode) {
	shortcodes[name] = shortcode
}
