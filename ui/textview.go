package ui

import "github.com/roblillack/spot"

type TextView struct {
	X        int
	Y        int
	Width    int
	Height   int
	Text     string
	FontSize int
	ref      nativeTypeTextView
}

var _ spot.Component = &TextView{}
var _ spot.Mountable = &TextView{}

func (c *TextView) Render(ctx *spot.RenderContext) spot.Component {
	return c
}
