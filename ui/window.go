package ui

import "github.com/roblillack/spot"

var _ spot.Control = &Window{}

func (c *Window) Render(ctx *spot.RenderContext) spot.Component {
	return c
}
