package ui

import "github.com/roblillack/spot"

type Dial struct {
	X              int
	Y              int
	Width          int
	Height         int
	Min            float64
	Max            float64
	Value          float64
	OnValueChanged func(float64)
	ref            nativeTypeDial
}

var _ spot.Component = &Dial{}
var _ spot.Mountable = &Dial{}

func (c *Dial) Render(ctx *spot.RenderContext) spot.Component {
	return c
}
