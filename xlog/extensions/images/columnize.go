package images

import (
	"github.com/emad-elsaid/xlog/markdown/ast"
	"github.com/emad-elsaid/xlog/markdown/parser"
	"github.com/emad-elsaid/xlog/markdown/text"
	"github.com/pnppl/wikilink"
)

type columnizeImagesParagraph struct{}

func (t columnizeImagesParagraph) Transform(doc *ast.Document, reader text.Reader, pc parser.Context) {
	paragraphs := []*ast.Paragraph{}

	ast.Walk(doc, func(node ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering {
			return ast.WalkContinue, nil
		}

		for c := node.FirstChild(); c != nil; c = c.NextSibling() {
			n, ok := c.(*ast.Paragraph)
			if !ok {
				continue
			}

			if containsOnlyImages(n) {
				paragraphs = append(paragraphs, n)
			}
		}

		return ast.WalkContinue, nil
	})

	for _, p := range paragraphs {
		removeBreaks(p)
		replaceWithColumns(p)
	}
}

func containsOnlyImages(n *ast.Paragraph) bool {
	if n.ChildCount() < 2 {
		return false
	}

	for c := n.FirstChild(); c != nil; c = c.NextSibling() {
		if c.Kind() != ast.KindImage && c.Kind() != ast.KindText && c.Kind() != wikilink.Kind {
			return false
		} else if wl, ok := c.(*wikilink.Node); ok {
		// if node is type wikilink, we only columnize if it's an image (Embed is true)
			if wl.Embed == false {
				return false
			}
		} else if t, ok := c.(*ast.Text); ok && !t.SoftLineBreak() {
			return false
		}
	}

	return true
}

func removeBreaks(n *ast.Paragraph) {
	breaks := []*ast.Text{}

	for c := n.FirstChild(); c != nil; c = c.NextSibling() {
		if t, ok := c.(*ast.Text); ok {
			breaks = append(breaks, t)
		}
	}

	for _, b := range breaks {
		n.RemoveChild(n, b)
	}
}

func replaceWithColumns(n *ast.Paragraph) {
	p := n.Parent()
	p.ReplaceChild(p, n, &imagesColumns{*n})
}
