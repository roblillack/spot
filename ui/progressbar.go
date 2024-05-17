package ui

import "github.com/roblillack/spot"

type ProgressBar struct {
	X             int
	Y             int
	Width         int
	Height        int
	Min           float64
	Max           float64
	Value         float64
	Indeterminate bool
	ref           nativeTypeProgressBar
}

var _ spot.Component = &ProgressBar{}
var _ spot.Control = &ProgressBar{}

func (c *ProgressBar) Render(ctx *spot.RenderContext) spot.Component {
	return c
}
