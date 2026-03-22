// Package ast defines AST nodes that represents extension's elements
package ast

import (
	gast "github.com/emad-elsaid/xlog/markdown/ast"
)

// A Typography struct represents a typographic entity
type Typography struct {
	gast.BaseInline

	Name string
	Sub string
	Orig []byte
}

// Dump implements Node.Dump.
func (n *Typography) Dump(source []byte, level int) {
	gast.DumpHelper(n, source, level, nil, nil)
}

// KindTypography is a NodeKind of the Typography node.
var KindTypography = gast.NewNodeKind("Typography")

// Kind implements Node.Kind.
func (n *Typography) Kind() gast.NodeKind {
	return KindTypography
}

// NewTypography returns a new Typography node.
//func NewTypography(name string, sub string) *Typography {
func NewTypography(name string, sub string, orig []byte) *Typography {
	return &Typography{
		Name: name,
		Sub: sub,
		Orig: orig,
	}
}
