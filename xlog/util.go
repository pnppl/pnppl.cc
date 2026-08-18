package xlog
import (
	"strings"
	"slices"
)

func GetTitle(page Page) string {
	content := page.Content()
	lines := strings.Split(string(content), "\n")
	firstLine := strings.TrimSpace(lines[0])
	normalizedName := strings.Replace(strings.Replace(firstLine, "# ", "", 1), " #", "", 1)
	if len([]byte(normalizedName)) <= 0 {
		return page.Name()
	}
	return normalizedName
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

