package opengraph

import (
	"flag"
	"fmt"
	"html/template"
	"net/url"
	"strings"

	. "github.com/emad-elsaid/xlog"

	"github.com/emad-elsaid/xlog/markdown/ast"
)

var domain string
var twitterUsername string

const descriptionLength = 200

func init() {
	flag.StringVar(&domain, "og.domain", "", "opengraph domain name to be used for meta tags of og:* and twitter:*")
	flag.StringVar(&twitterUsername, "twitter.username", "", "user twitter account @handle. including the @")

	RegisterExtension(Opengraph{})
}

type Opengraph struct{}

func (Opengraph) Name() string { return "opengraph" }
func (Opengraph) Init() {
	RegisterWidget(WidgetHead, 1, opengraphTags)
}

func opengraphTags(p Page) template.HTML {
	escape := template.JSEscapeString

	name := p.Name()
	if p.Name() == Config.Index {
		name = Config.Sitename
	}

	props := Properties(p)
	var title string
	if v, ok := props["title"]; ok && v != nil {
		if s, ok := v.Value().(string); ok {
			title = s
		} else {
			title = name
		}
	} else {
		title = name
	}

	var u url.URL
	u.Scheme = "https"
	u.Host = domain
	u.Path = "/" + name

	URL := u.String()

	var image string
	src, tree := p.AST()
	if imageAST, ok := FindInAST[*ast.Image](tree); ok && imageAST != nil {
		image = "https://" + domain + string(imageAST.Destination)
	} else {
		image = "https://" + domain + "/public/logo.gif"
	}

	firstParagraph := escape(rawText(src, tree, descriptionLength))
//	The escape fails miserably
	firstParagraph = strings.ReplaceAll(firstParagraph, "\\\"", "&quot;")
	firstParagraph = strings.ReplaceAll(firstParagraph, "\\'", "&apos;")

	ogTags := fmt.Sprintf(`
		<meta property="og:site_name" content="%s">
		<meta property="og:title" content="%s">
		<meta property="og:description" content="%s">
		<meta property="og:image" content="%s">
		<meta property="og:url" content="%s">
		<meta property="og:type" content="website">
`,
		escape(Config.Sitename),
		escape(title),
		firstParagraph,
		escape(image),
		escape(URL),
	)

// Fuck xitter
/*	twitterTags := fmt.Sprintf(`
    <meta name="twitter:title" content="%s" />
    <meta name="twitter:description" content="%s" />
    <meta name="twitter:image" content="%s" />
    <meta name="twitter:card" content="summary_large_image" />
    <meta name="twitter:creator" content="%s" />
    <meta name="twitter:site" content="%s" />
    <meta name="twitter:image:alt" content="%s" />
`,
		escape(title),
		escape(firstParagraph),
		escape(image),
		escape(twitterUsername),
		escape(twitterUsername),
		escape(title),
	)
*/
	metaTags := fmt.Sprintf(`		<meta name="description" content="%s">`,
		firstParagraph,
	)

	return template.HTML(ogTags + metaTags)
}

func rawText(source []byte, n ast.Node, limit int) string {
	if source == nil || n == nil {
		return ""
	}

	out := ""
	ast.Walk(n, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}

		if n.Kind() == ast.KindText {
			out += " " + strings.TrimSpace(string(n.(*ast.Text).Text(source)))
		}

		if len(out) > limit {
			out = out[:limit]
			return ast.WalkStop, nil
		}

		return ast.WalkContinue, nil
	})

	return strings.TrimSpace(out)
}
