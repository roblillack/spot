//go:build !fltk && (darwin || cocoa)

package ui

import (
	"github.com/roblillack/gocoa"
	"github.com/roblillack/spot"
)

type nativeTypeSlider = *gocoa.Slider

func (b *Slider) Update(nextComponent spot.Control) bool {
	next, ok := nextComponent.(*Slider)
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

	b.OnValueChanged = next.OnValueChanged

	return true
}

func (b *Slider) Mount(parent spot.Control) any {
	if b.ref != nil {
		return b.ref
	}

	b.ref = gocoa.NewSlider(b.X, b.Y, b.Width, b.Height)
	b.ref.SetMaximumValue(b.Max)
	b.ref.SetMinimumValue(b.Min)
	b.ref.SetValue(b.Value)
	b.ref.OnSliderValueChanged(b.callback)

	if window, ok := parent.(*Window); ok && window != nil && window.ref != nil {
		window.ref.AddSlider(b.ref)
	}

	return b.ref
}

func (c *Slider) Unmount() {
	if c.ref == nil {
		return
	}

	c.ref.Remove()
	c.ref = nil
}

func (c *Slider) callback() {
	if c.ref == nil {
		return
	}

	if c.OnValueChanged != nil {
		c.OnValueChanged(c.ref.Value())
	}
}
