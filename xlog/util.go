package xlog
import (
	"strings"
	"slices"
	"errors"
	"unicode/utf8"
	"fmt"

	"github.com/emad-elsaid/xlog/markdown/util"
	"git.sr.ht/~timetoplatypus/recfile"
)

func DecimalEscape(input string) string {
	source := []byte(input)
	output := []byte("")
	for i := 0; i < len(source); {
		b := source[i]
		if b < utf8.RuneSelf {
			if v := util.EscapeHTMLByte(b); v != nil {
				output = append(output, v...)
			} else {
				output = append(output, b)
			}
			i++
		} else {
			r, size := utf8.DecodeRune(source[i:])
			output = append(output, fmt.Sprintf("&#%d;", r)...)
			i += size
		}
	}
	return string(output)
}

func NormalizeName(name string) (string, error) {
	firstLine := strings.TrimSpace(name)
	normalizedName := strings.TrimSuffix(strings.TrimPrefix(firstLine, "# "), " #")
	if len([]byte(normalizedName)) <= 0 {
		return name, errors.New("Couldn't normalize name")
	}
	return normalizedName, nil
}

func GetTitle(page Page) string {
	content := page.Content()
	lines := strings.Split(string(content), "\n")
	name, err := NormalizeName(lines[0])
	if err != nil {
		return page.Name()
	}
	return name
}

func PageTitleCompare(a, b Page) int {
	return strings.Compare(strings.ToLower(GetTitle(a)), strings.ToLower(GetTitle(b)))
}

func PageTitleSort(pages []Page) []Page {
	pages = slices.DeleteFunc(pages, func(a Page) bool {
		return IsIgnoredPath(a.Name())
	})

	slices.SortFunc(pages, func(a, b Page) int {
		return PageTitleCompare(a, b)
	})
	return pages
}

// rectype == "0" for unspecified
func GetFromRec(rectype string, keyname string, keyval string, fieldname string, rec recfile.Database) ([]string, error) {
	var results []string
	for _, set := range rec.RecordSets {
		if string(set.Descriptor.Type) == rectype || rectype == "0" {
			for _, record := range set.Records {
				if slices.Contains(record.Fields, recfile.Field{Name: recfile.Name(keyname), Value: recfile.Value(keyval)}) {
					for _, field := range record.Fields {
						if string(field.Name) == fieldname {
							results = append(results, string(field.Value))
						}
					}
				}
			}
		}
		if rectype == "0" {
			break
		}
	}
	if len(results) > 0 {
		return results, nil
	}
	return nil, errors.New("No matches")
}
