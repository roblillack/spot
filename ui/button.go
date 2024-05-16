package ui

import "github.com/roblillack/spot"

type Button struct {
	X       int
	Y       int
	Width   int
	Height  int
	Title   string
	OnClick func()

	ref nativeTypeButton
}

var _ spot.Component = &Button{}
var _ spot.Control = &Button{}

func (b *Button) Render(ctx *spot.RenderContext) spot.Component {
	return b
}
