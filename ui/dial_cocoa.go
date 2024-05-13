//go:build !fltk && (darwin || cocoa)

package ui

import (
	"github.com/mojbro/gocoa"
	"github.com/roblillack/spot"
)

type Dial struct {
	X              int
	Y              int
	Width          int
	Height         int
	Min            float64
	Max            float64
	Value          float64
	OnValueChanged func(float64)
	ref            *gocoa.Slider
}

func (b *Dial) Equals(other spot.Component) bool {
	next, ok := other.(*Dial)
	if !ok {
		return false
	}

	if b == nil && next != nil || b != nil && next == nil {
		return false
	}

	return next.Max == b.Max && next.Min == b.Min &&
		next.Value == b.Value
}

func (b *Dial) Update(nextComponent spot.Component) bool {
	next, ok := nextComponent.(*Dial)
	if !ok {
		return false
	}

	if b.ref == nil {
		return false
	}

	if next.Min != b.Min {
		b.Min = next.Min
		b.ref.SetMinimumValue(b.Min)
	}

	if next.Max != b.Max {
		b.Max = next.Max
		b.ref.SetMaximumValue(b.Max)
	}

	if next.Value != b.Value {
		b.Value = next.Value
		b.ref.SetValue(b.Value)
	}

	return true
}

func (b *Dial) Mount() any {
	if b.ref != nil {
		return b.ref
	}

	b.ref = gocoa.NewSlider(b.X, b.Y, b.Width, b.Height)
	b.ref.SetMaximumValue(b.Max)
	b.ref.SetMinimumValue(b.Min)
	b.ref.SetValue(b.Value)
	b.ref.SetSliderType(gocoa.SliderTypeCircular)
	b.ref.OnSliderValueChanged(func() {
		if b.OnValueChanged != nil {
			b.OnValueChanged(b.ref.Value())
		}
	})
	return b.ref
}

var _ spot.Component = &Dial{}
