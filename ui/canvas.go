package ui

import (
	"github.com/roblillack/spot"
	"github.com/roblillack/spot/ui/canvas"
)

type Canvas struct {
	X        int
	Y        int
	Width    int
	Height   int
	OnClick  func(x, y int)
	Elements []canvas.Element

	ref nativeTypeCanvas
}

var _ spot.Component = &Canvas{}
var _ spot.Control = &Canvas{}

func (b *Canvas) Render(ctx *spot.RenderContext) spot.Component {
	return b
}
