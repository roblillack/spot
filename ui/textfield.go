package ui

import (
	"github.com/roblillack/spot"
)

type TextField struct {
	X      int
	Y      int
	Width  int
	Height int
	Value  string
	// FontSize int
	OnChange func(value string)
	// Editable     bool
	// Bezeled      bool
	// Selectable   bool
	// NoBackground bool

	ref nativeTypeTextField
}

var _ spot.Component = &TextField{}
var _ spot.Mountable = &TextField{}

func (c *TextField) Render(ctx *spot.RenderContext) spot.Component {
	return c
}
