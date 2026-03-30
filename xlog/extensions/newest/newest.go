package newest

import (
	"fmt"
	"html/template"
	"strings"
	"regexp"

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
		return template.HTML(fmt.Sprintf("[Placeholder. Refresh the page.]"))
	}
	p := xlog.NewPage(newestName)
	props := xlog.Properties(p)
	title := newestName
	if v, ok := props["title"]; ok && v != nil {
		if s, ok := v.Value().(string); ok {
			title = s
		}
	}

	// get last heading to jump to right spot
    pattern := regexp.MustCompile(`<h[1-6] id="(?P<id>.+?)">`)
    var result []byte
    var ids []string
	var content []string
	// do we actually know if it's safe to break the html up by line?
	// answer: it is not...
	i := 0
	for line := range strings.Lines(string(p.Render())) {
		if i == 5 { break }
		content = append(content, line)
		// heading id
		for _, submatches := range pattern.FindAllSubmatchIndex([]byte(line), -1) {
			ids = append(ids, string(pattern.ExpandString(result, "$id", line, submatches)))
		}
		i++
	}

	id := ""
	if len(ids) > 0 {
		id = "#" + ids[len(ids)-1]
	}

	contents := strings.TrimSpace(strings.Join(content[0:], ""))
	// hideous quick fix for relative paths issue
	pathfix := regexp.MustCompile(`src="\.\.`)
	contents = pathfix.ReplaceAllString(contents, `src="`)
	pathfix = regexp.MustCompile(`href="\.\.`)
	contents = pathfix.ReplaceAllString(contents, `href="`)
	return template.HTML(fmt.Sprintf(`<h2 class="excerpt">Newest post: <a href="%s" class="excerpt">%s</a></h2> <blockquote class="excerpt">%s<a href="%s" class="excerpt-more">...</a></blockquote>`, newestName, title, contents, newestName + id))
}
