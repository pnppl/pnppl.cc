package hashtags

import (
	"fmt"
	"slices"
	"github.com/emad-elsaid/xlog"
)

type Prop struct {
	NameVal string
	Val     []string
}

func (p Prop) Name() string { return p.NameVal }
func (p Prop) Value() any   { return p.Val }
func (p Prop) Icon() string { return "" }
func properties(p xlog.Page) []xlog.Property {
	props := []xlog.Property{}
	switch t := p.(type) {
		default:
			fmt.Printf("Type error in hashtags/properties.go. Type: %T", t)
			return props
		case xlog.DynamicPage:
			return props
		case xlog.Page:
			hashtags := h.hashtagsFor(p)
			var tags []string
			for _, v := range(hashtags) {
				tags = append(tags, string(v.value))
			}
			slices.Sort(tags)
			props = append(props, Prop{
					NameVal: "tags",
					Val:	tags,
				})
			return props
	}
}
