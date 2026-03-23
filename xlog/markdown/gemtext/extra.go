package gemtext

import (
	"fmt"

	"git.sr.ht/~kota/fuckery"
	"github.com/emad-elsaid/xlog/markdown/ast"
	east "github.com/emad-elsaid/xlog/markdown/extension/ast"
	"github.com/emad-elsaid/xlog/markdown/util"
)

// renderStrikethrough writes strikethrough text based on a few config options.
func (r *GemRenderer) renderStrikethrough(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	n := node.(*east.Strikethrough)
	if entering {
		switch r.config.Strikethrough {
		case StrikethroughMarkdown:
			fmt.Fprintf(w, "~~")
		case StrikethroughUnicode:
			fmt.Fprintf(w, "%s", fuckery.Strike(string(n.Text(source))))
			return ast.WalkSkipChildren, nil
		}
	} else {
		switch r.config.Strikethrough {
		case StrikethroughMarkdown:
			fmt.Fprintf(w, "~~")
		}
	}
	return ast.WalkContinue, nil
}

// renderWiki writes a wiki style link in gemtext.
// Similar to links and autolinks the node is skipped if the parent node
// contains only links. We use the parent node (paragraph or heading) to do the
// actual heavy lifting.
func (r *GemRenderer) renderWiki(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	// skip if the parent node contains only links
	if linkOnly(source, node.Parent()) {
		return ast.WalkSkipChildren, nil
	}
	curly := r.config.ParagraphLink == ParagraphLinkCurlyBelow
	if entering {
		if curly {
			fmt.Fprint(w, "{")
		}
	} else {
		if curly {
			fmt.Fprint(w, "}")
		}
	}

	return ast.WalkContinue, nil
}

// renderTypography renders typographic entity nodes
func (r *GemRenderer) renderTypography(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering && node.Kind() == east.KindTypography {
		n := node.(*east.Typography)
		fmt.Fprintf(w, string(n.Orig))
	}
	return ast.WalkContinue, nil
}


// renderFootnoteLink renders the superscript number in the main text
func (r *GemRenderer) renderFootnoteLink(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering && node.Kind() == east.KindFootnoteLink {
		n := node.(*east.FootnoteLink)
		number := n.Index
		// this is pretty but impractical, can't jump back and forth with search
		/*
		var sups []string
		// gotta get each place one at a time
		for i := number; i > 0; i /= 10 {
			sups = append(getSup(i % 10), sups...)
			i -= i % 10
		}
		for i := range sups {
			fmt.Fprintf(w, sups[i])
			// zero-width nbsp
			if i < len(sups) - 1 {
				fmt.Fprintf(w, "\u2060")
			}
		}
		*/
		fmt.Fprintf(w, "[#%d]", number)
	}
	return ast.WalkContinue, nil
}

/*
func getSup(num int) []string {
	var sup string
	switch num {
		case 0:
			sup = "⁰"
		case 1:
			sup = "¹"
		case 2:
			sup = "²"
		case 3:
			sup = "³"
		case 4:
			sup = "⁴"
		case 5:
			sup = "⁵"
		case 6:
			sup = "⁶"
		case 7:
			sup = "⁷"
		case 8:
			sup = "⁸"
		case 9:
			sup = "⁹"
	}
	return []string{sup}
}
*/

func (r *GemRenderer) renderFootnoteList(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if entering && node.Kind() == east.KindFootnoteList {
		fmt.Fprintf(w, "# Footnotes\n")
	}
	return ast.WalkContinue, nil
}

func (r *GemRenderer) renderFootnote(w util.BufWriter, source []byte, node ast.Node, entering bool) (ast.WalkStatus, error) {
	if node.Kind() == east.KindFootnote {
		n := node.(*east.Footnote)
		if entering {
			fmt.Fprintf(w, "## [#%d]:\n", n.Index)
		}
	}
	return ast.WalkContinue, nil
}
