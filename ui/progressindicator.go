package ui

import "github.com/roblillack/spot"

type ProgressIndicator struct {
	X             int
	Y             int
	Width         int
	Height        int
	Min           float64
	Max           float64
	Value         float64
	Indeterminate bool
	ref           nativeTypeProgressIndicator
}

var _ spot.Component = &ProgressIndicator{}
var _ spot.Control = &ProgressIndicator{}

func (c *ProgressIndicator) Render(ctx *spot.RenderContext) spot.Component {
	return c
}
