package autolink_pages

import (
	"strings"

	. "github.com/emad-elsaid/xlog"
	"github.com/emad-elsaid/xlog/markdown/ast"
	"github.com/emad-elsaid/xlog/markdown/parser"
	"github.com/emad-elsaid/xlog/markdown/text"
	"github.com/emad-elsaid/xlog/markdown/util"
)

type pageLinkParser struct{}

func (*pageLinkParser) Trigger() []byte {
	// ' ' indicates any white spaces and a line head
//	return []byte{' ', '*', '_', '~', '(', '!', '@', '&', '|', ';', '?'}
	return []byte{' ', '*', '_', '~', '('}
}

// getCurrentPageTitle extracts the normalized title from the document source
func getCurrentPageTitle(source []byte) string {
	lines := strings.SplitN(string(source), "\n", 2)
	if len(lines) == 0 {
		return ""
	}
	firstLine := strings.TrimSpace(lines[0])
	firstLine = strings.TrimPrefix(firstLine, "# ")
	firstLine = strings.TrimSuffix(firstLine, " #")
	return strings.ToLower(firstLine)
}



func (s *pageLinkParser) Parse(parent ast.Node, block text.Reader, pc parser.Context) ast.Node {
	if pc.IsInLinkLabel() {
		return nil
	}

	if autolinkPages == nil {
		UpdatePagesList(nil)
	}

	line, segment := block.PeekLine()
	if line == nil {
		return nil
	}

	consumes := 0
	start := segment.Start
	c := line[0]
	// advance if current position is not a line head.
	if c == ' ' || c == '*' || c == '_' || c == '~' || c == '(' {
		consumes++
		start++
		line = line[1:]
	}

	var found Page
	var m int
	normalizedLine := strings.ToLower(string(line))

	// Get current page title to avoid self-linking
	currentPageTitle := getCurrentPageTitle(block.Source())

	for _, p := range autolinkPages {
		if len(line) < len(p.normalizedName) {
			continue
		}

		// Skip if this would link to the current page
		if p.normalizedName == currentPageTitle {
			continue
		}

		// Found a page
		if strings.HasPrefix(normalizedLine, p.normalizedName) {
			found = p.page
			m = len(p.normalizedName)
			break
		}
	}

	if found == nil ||
		(len(line) > m && util.IsAlphaNumeric(line[m])) { // next character is word character
		block.Advance(consumes)
		return nil
	}

	if consumes != 0 {
		s := segment.WithStop(segment.Start + 1)
		ast.MergeOrAppendTextSegment(parent, s)
	}
	consumes += m
	block.Advance(consumes)

	n := ast.NewTextSegment(text.NewSegment(start, start+m))
	link := &PageLink{
		page: found,
	}
	link.AppendChild(link, n)
	return link
}
