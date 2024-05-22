package ui

import (
	"github.com/roblillack/spot"
)

type LabelAlignment int

const (
	LabelAlignmentLeft LabelAlignment = iota
	LabelAlignmentCenter
	LabelAlignmentRight
)

type Label struct {
	X        int
	Y        int
	Width    int
	Height   int
	Value    string
	FontSize int
	Align    LabelAlignment
	ref      nativeTypeLabel
}

var _ spot.Component = &Label{}
var _ spot.Control = &Label{}

func (c *Label) Render(ctx *spot.RenderContext) spot.Component {
	return c
}
