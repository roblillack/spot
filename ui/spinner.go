//go:build ignore

package ui

import "github.com/roblillack/spot"

type Spinner struct {
	X              int
	Y              int
	Width          int
	Height         int
	Min            float64
	Max            float64
	Step           float64
	Value          float64
	OnValueChanged func(float64)
	ref            nativeTypeSpinner
}

var _ spot.Component = &Spinner{}
var _ spot.Control = &Spinner{}

func (c *Spinner) Render(ctx *spot.RenderContext) spot.Component {
	return c
}
