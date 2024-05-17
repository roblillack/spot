package ui

import (
	"github.com/roblillack/spot"
)

type Dropdown struct {
	X                    int
	Y                    int
	Width                int
	Height               int
	Items                []string
	SelectedIndex        int
	Editable             bool
	OnSelectionDidChange func(index int)
	ref                  nativeTypeDropdown
}

var _ spot.Component = &Dropdown{}
var _ spot.Control = &Dropdown{}

func (c *Dropdown) Render(ctx *spot.RenderContext) spot.Component {
	return c
}
