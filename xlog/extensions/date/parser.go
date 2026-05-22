package date

import (
	"time"
	"unicode"

	"github.com/emad-elsaid/xlog"
	"github.com/emad-elsaid/xlog/markdown/ast"
	"github.com/emad-elsaid/xlog/markdown/parser"
	"github.com/emad-elsaid/xlog/markdown/text"
)

type dateParser struct{}

func (s *dateParser) Trigger() []byte {
	return []byte{' '}
}

type dateFilenameParser struct{}

func (s *dateFilenameParser) Trigger() []byte {
	return []byte{' '}
}

var (
	datePatterns = []string{
		`2006-01-02`,
//		`2006-1-2`,

//		`2006-January-2`,
//		`2006/January/2`,
//		`2006\January\2`,

//		`2006-Jan-2`,
//		`2006/Jan/2`,
//		`2006\Jan\2`,

//		`2-January-2006`,
//		`2/January/2006`,
//		`2\January\2006`,

//		`2-Jan-2006`,
//		`2/Jan/2006`,
//		`2\Jan\2006`,

//		`Jan-2-2006`,
//		`Jan/2/2006`,
//		`Jan\2\2006`,

//		`January-2-2006`,
//		`January/2/2006`,
//		`January\2\2006`,

		`01/02/2006`,
		`01/02/06`,
		`1/2/2006`,
		`1/2/06`,
//		`January 2, 2006`,
//		`January 02, 2006`,
//		`Jan 2, 2006`,
//		`Jan 02, 2006`,
//		`Jan. 2, 2006`,
//		`Jan. 02, 2006`,
//		`Jan. 2nd, 2006`,
//		`Jan 2nd, 2006`,
//		`January 2nd, 2006`,
	}
)

var filenameProcessedKey = parser.NewContextKey()
func (s *dateFilenameParser) Parse(parent ast.Node, reader text.Reader, pc parser.Context) ast.Node {
	if pc.Get(filenameProcessedKey) == nil {
		pc.Set(filenameProcessedKey, true)
		if filename, ok := pc.Get(xlog.PageFilenameKey).(string); ok {
			if len(filename) >= 10 {
				t, err := time.Parse("2006-01-02", filename[:10])
				if err == nil {
					return &DateNode{
						time: t,
						drawTag: false,
					}
				}
			}
		}
	}
	return nil
}

func (s *dateParser) Parse(parent ast.Node, reader text.Reader, pc parser.Context) ast.Node {
	if filename, ok := pc.Get(xlog.PageFilenameKey).(string); ok {
		if filename == "deathnote.md" {
			return nil
		}
	}

	l, _ := reader.PeekLine()
	if len(l) < 2 {
		return nil
	}

	advance := 0

	if l[0] == ' ' {
		advance++
		l = l[1:]
	}

	space := len(l)
	separators := 0
	for i, b := range l {
		if !unicode.In(rune(b), unicode.Digit, unicode.Letter, unicode.Dash) &&
			b != '/' &&
			b != '\\' {
			space = i
			break
		}

		// keep track of how many separators
		if unicode.In(rune(b), unicode.Dash) || b == '/' || b == '\\' {
			separators++
		}

		if separators > 2 {
			space = i
			break
		}
	}

	advance += space
	l = l[:space]

	for _, pattern := range datePatterns {
		t, err := time.Parse(pattern, string(l))
		if err == nil {
			reader.Advance(advance)
			return &DateNode{
				time: t,
				drawTag: true,
			}
		}
	}

	return nil
}
