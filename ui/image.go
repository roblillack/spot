package ui

import (
	"image"

	"github.com/roblillack/spot"
)

type Image struct {
	X       int
	Y       int
	Width   int
	Height  int
	Border  bool
	OnClick func(x, y int, secondary bool)
	Image   image.Image

	ref nativeTypeImage
}

var _ spot.Component = &Image{}
var _ spot.Mountable = &Image{}

func (c *Image) Render(ctx *spot.RenderContext) spot.Component {
	return c
}
