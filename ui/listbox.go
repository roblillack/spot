package ui

import (
	"github.com/roblillack/spot"
)

type ListBox struct {
	X           int
	Y           int
	Width       int
	Height      int
	Values      []string
	Multiselect bool
	Selection   []int
	OnSelect    func([]int)
	ref         nativeTypeListBox
}

var _ spot.Component = &ListBox{}
var _ spot.Control = &ListBox{}

func (c *ListBox) Render(ctx *spot.RenderContext) spot.Component {
	return c
}
