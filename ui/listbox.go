package ui

import (
	"github.com/roblillack/spot"
)

type Listbox struct {
	X           int
	Y           int
	Width       int
	Height      int
	Values      []string
	Multiselect bool
	Selection   []int
	OnSelect    func([]int)
	ref         nativeTypeListbox
}

var _ spot.Component = &Listbox{}
var _ spot.Control = &Listbox{}

func (c *Listbox) Render(ctx *spot.RenderContext) spot.Component {
	return c
}
