package ui

import "github.com/roblillack/spot"

type Slider struct {
	X      int
	Y      int
	Width  int
	Height int
	Min    float64
	Max    float64
	Value  float64
	// Type           goFltk.SliderType
	// Type           gocoa.SliderType
	OnValueChanged func(float64)
	ref            nativeTypeSlider
}

var _ spot.Component = &Slider{}
var _ spot.Control = &Slider{}

func (c *Slider) Render(ctx *spot.RenderContext) spot.Component {
	return c
}
