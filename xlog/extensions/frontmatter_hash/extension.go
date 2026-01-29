package frontmatter_hash

import (
	"github.com/emad-elsaid/xlog"
)

func init() {
	xlog.RegisterExtension(Frontmatter_hash{})
}

type Frontmatter_hash struct{}

func (Frontmatter_hash) Name() string { return "frontmatter_hash" }
func (Frontmatter_hash) Init() {
	m := New(
		WithStoresInDocument(),
	)

	m.Extend(xlog.MarkdownConverter())
	xlog.RegisterProperty(MetaProperties)
}

type MetaProperty struct {
	NameVal string
	Val     any
}

func (m MetaProperty) Name() string { return m.NameVal }
func (m MetaProperty) Icon() string { return "fa-solid fa-table-list" }
func (m MetaProperty) Value() any   { return m.Val }

func MetaProperties(p xlog.Page) []xlog.Property {
	_, ast := p.AST()
	if ast == nil {
		return nil
	}

	metaData := ast.OwnerDocument().Meta()
	if len(metaData) == 0 {
		return nil
	}

	ps := make([]xlog.Property, 0, len(metaData))
	for k, v := range metaData {
		ps = append(ps, MetaProperty{
			NameVal: k,
			Val:     v,
		})
	}

	return ps
}
