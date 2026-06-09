package extension

import (
//	"unicode/utf8"
	"fmt"
	"github.com/emad-elsaid/xlog/markdown"
	gast "github.com/emad-elsaid/xlog/markdown/ast"
	"github.com/emad-elsaid/xlog/markdown/extension/ast"
	"github.com/emad-elsaid/xlog/markdown/renderer"
	"github.com/emad-elsaid/xlog/markdown/renderer/html"
	"github.com/emad-elsaid/xlog/markdown/parser"
	"github.com/emad-elsaid/xlog/markdown/text"
	"github.com/emad-elsaid/xlog/markdown/util"
)

var uncloseCounterKey = parser.NewContextKey()

type unclosedCounter struct {
	Single int
	Double int
}

func (u *unclosedCounter) Reset() {
	u.Single = 0
	u.Double = 0
}

func getUnclosedCounter(pc parser.Context) *unclosedCounter {
	v := pc.Get(uncloseCounterKey)
	if v == nil {
		v = &unclosedCounter{}
		pc.Set(uncloseCounterKey, v)
	}
	return v.(*unclosedCounter)
}

func subMap() map[string]string {
	return map[string]string {
		"ndash": "--",
		"mdash": "---",
		"hellip": "...",
		"lsquo": "'",
		"rsquo": "'",
		"ldquo": "\"",
		"rdquo": "\"",
//		"laquo": "<<",
//		"raquo": ">>",
//		"iquest": "?",
//		"iexcl": "!",
		"ne": "!=",
		"asymp": "~=",
		"ge": ">=",
		"le": "<=",
//		"ntilde": "n",
//		"Ntilde": "N",
		"times": "x",
	}
}

//type typographerDelimiterProcessor struct {
//}
//
//func (p *typographerDelimiterProcessor) IsDelimiter(b byte) bool {
//	return b == '\'' || b == '"'
//}
//
//func (p *typographerDelimiterProcessor) CanOpenCloser(opener, closer *parser.Delimiter) bool {
//	return opener.Char == closer.Char
//}
//
//func (p *typographerDelimiterProcessor) OnMatch(consumes int) gast.Node {
//	return nil
//}
//
//var defaultTypographerDelimiterProcessor = &typographerDelimiterProcessor{}

type typographerParser struct {
}

// NewTypographerParser return a new InlineParser that parses
// typographer expressions.
func NewTypographerParser() parser.InlineParser {
	return &typographerParser{}
}

func (s *typographerParser) Trigger() []byte {
//	return []byte{'\'', '"', '-', '.', ',', '<', '>', '*', '['}
	entityByte := []byte(" ")
	for name, _ := range subMap() {
		entity, ok := util.LookUpHTML5EntityByName(name)
		if ok {
			char := entity.Characters[0]
			entityByte = append(entityByte, char)
//			fmt.Printf("%d   ", entity.Characters)
		}
	}
//	return []byte{0xE2, 0xC2, ' '}
	return entityByte
}

func spanner(name string, sub string) string {
	return fmt.Sprintf("%s%s%s%s%s", `<span class="`, name, `"></span><span class="is-hidden">`, sub, `</span>`)
}

func (s *typographerParser) Parse(parent gast.Node, block text.Reader, pc parser.Context) gast.Node {
	line, _ := block.PeekLine()
	if len(line) > 1 {
		for name, sub := range subMap() {
			entity, ok := util.LookUpHTML5EntityByName(name)
				if ok {
					bytes := entity.Characters
					if line[0] == bytes[0] && line[1] == bytes[1] {
						if len(bytes) == 2 {
//							node := gast.NewString(spanner(name, sub))
//							node.SetCode(true)
							node := ast.NewTypography(name, sub, line[:2])
							block.Advance(2)
							return node
						}
						if len(line) > 2 {
							if line[2] == bytes[2] {
//								node := gast.NewString(spanner(name, sub))
//								node.SetCode(true)
								node := ast.NewTypography(name, sub, line[:3])
								block.Advance(3)
								return node
							}
						}
					}
				}
		}
	}
	return nil
}

func (s *typographerParser) CloseBlock(parent gast.Node, pc parser.Context) {
	getUnclosedCounter(pc).Reset()
}

type typographer struct {
}

// Typographer is an extension that replaces punctuations with typographic entities.
var Typographer = &typographer{}

// NewTypographer returns a new Extender that replaces punctuations with typographic entities.
func NewTypographer() markdown.Extender {
	return &typographer{
	}
}

// TypographyHTMLRenderer is a renderer.NodeRenderer implementation that
// renders Typography nodes.
type TypographyHTMLRenderer struct {
}

// TypographyHTMLRenderer returns a new TypographyHTMLRenderer.
func NewTypographyHTMLRenderer(opts ...html.Option) renderer.NodeRenderer {
	return &TypographyHTMLRenderer{}
}

// RegisterFuncs implements renderer.NodeRenderer.RegisterFuncs.
func (r *TypographyHTMLRenderer) RegisterFuncs(reg renderer.NodeRendererFuncRegisterer) {
	reg.Register(ast.KindTypography, r.renderTypography)
}

func (r *TypographyHTMLRenderer) renderTypography(
	w util.BufWriter, source []byte, n gast.Node, entering bool) (gast.WalkStatus, error) {
	if entering && n.Kind() == ast.KindTypography {
		// type assertion needed for some dark wizard type system purposes i don't understand
		node := n.(*ast.Typography)
		_, _ = w.WriteString(spanner(node.Name, node.Sub))
	}
	return gast.WalkContinue, nil
}

func (e *typographer) Extend(m markdown.Markdown) {
	m.Parser().AddOptions(parser.WithInlineParsers(
//		util.Prioritized(NewTypographerParser(e.options...), 9999),
		util.Prioritized(NewTypographerParser(), 9999),
	))
	m.Renderer().AddOptions(renderer.WithNodeRenderers(
		util.Prioritized(NewTypographyHTMLRenderer(), 9999),
	))
}
