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
	name, err := NormalizeName(lines[0])
	if err != nil {
		return ""
	}
	return strings.ToLower(name)
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
	var consumedChars int
	normalizedLine := strings.ToLower(string(line))

	// Get current page title to avoid self-linking
	currentPageTitle := getCurrentPageTitle(block.Source())

	outer:
	for _, p := range autolinkPages {
		if len(line) < len(p.normalizedName) && len(p.aliases) == 0 {
			continue
		}

		// Skip if this would link to the current page
		if p.normalizedName == currentPageTitle {
			continue
		}

		// Found a page
		if strings.HasPrefix(normalizedLine, p.normalizedName) && !p.nodefault {
			found = p.page
			consumedChars = len(p.normalizedName)
			break
		}

		// page minus 'the'
		if strings.HasPrefix(normalizedLine, strings.TrimPrefix(p.normalizedName, "the ")) && !p.nodefault {
			found = p.page
			consumedChars = len(p.normalizedName) - 4
			break
		}

		// aliases
		for _, alias := range p.aliases {
			if strings.HasPrefix(normalizedLine, alias) {
				found = p.page
				consumedChars = len(alias)
				break outer
			}
		}
	}

	if found == nil ||
		(len(line) > consumedChars && util.IsAlphaNumeric(line[consumedChars])) { // next character is word character
		block.Advance(consumes)
		return nil
	}

	if consumes != 0 {
		s := segment.WithStop(segment.Start + 1)
		ast.MergeOrAppendTextSegment(parent, s)
	}
	consumes += consumedChars
	block.Advance(consumes)

	n := ast.NewTextSegment(text.NewSegment(start, start+consumedChars))
	link := &PageLink{
		page: found,
	}
	link.AppendChild(link, n)
	return link
}
