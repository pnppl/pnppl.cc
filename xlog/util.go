package xlog
import (
	"strings"
	"slices"
)
func PageTitleSort(pages []Page) []Page {
	pages = slices.DeleteFunc(pages, func(a Page) bool {
		return IsIgnoredPath(a.Name())
	})

	slices.SortFunc(pages, func(a, b Page) int {
		nameA := a.Name()
		nameB := b.Name()

		content := a.Content()
		lines := strings.Split(string(content), "\n")
		firstLine := strings.TrimSpace(lines[0])
		normalizedNameA := strings.Replace(strings.Replace(firstLine, "# ", "", 1), " #", "", 1)

		content = b.Content()
		lines = strings.Split(string(content), "\n")
		firstLine = strings.TrimSpace(lines[0])
		normalizedNameB := strings.Replace(strings.Replace(firstLine, "# ", "", 1), " #", "", 1)

		if len([]byte(normalizedNameA)) > 0 {
			nameA = normalizedNameA
		}
		if len([]byte(normalizedNameB)) > 0 {
			nameB = normalizedNameB
		}
		return strings.Compare(strings.ToLower(nameA), strings.ToLower(nameB))
	})
	return pages
}

