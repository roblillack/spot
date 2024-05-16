package ui

import (
	"github.com/roblillack/spot"
)

type ComboBox struct {
	X                    int
	Y                    int
	Width                int
	Height               int
	Items                []string
	SelectedIndex        int
	Editable             bool
	OnSelectionDidChange func(index int)
	ref                  nativeTypeComboBox
}

var _ spot.Component = &ComboBox{}
var _ spot.Control = &ComboBox{}

func (c *ComboBox) Render(ctx *spot.RenderContext) spot.Component {
	return c
}
