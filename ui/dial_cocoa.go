//go:build !fltk && (darwin || cocoa)

package ui

import (
	"github.com/roblillack/spot"
	"github.com/roblillack/spot/ui/internal/cocoa"
)

type nativeTypeDial = *cocoa.Slider

func (b *Dial) Update(nextComponent spot.Control) bool {
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

func (b *Dial) Mount(ctx *spot.RenderContext, parent spot.Control) any {
	if b.ref != nil {
		return b.ref
	}

	b.ref = cocoa.NewSlider(b.X, b.Y, b.Width, b.Height)
	b.ref.SetMaximumValue(b.Max)
	b.ref.SetMinimumValue(b.Min)
	b.ref.SetValue(b.Value)
	b.ref.SetSliderType(cocoa.SliderTypeCircular)
	b.ref.OnSliderValueChanged(func() {
		if b.OnValueChanged != nil {
			b.OnValueChanged(b.ref.Value())
		}
	})

	if window, ok := parent.(*Window); ok && window != nil && window.ref != nil {
		window.ref.AddSlider(b.ref)
	}

	return b.ref
}

func (c *Dial) Unmount() {
	if c.ref == nil {
		return
	}

	c.ref.Remove()
	c.ref = nil
}

func (c *Dial) Layout(ctx *spot.RenderContext, parent spot.Control) {
	x, y, w, h := calcLayout(parent, c.X, c.Y, c.Width, c.Height)
	c.ref.SetFrame(x, y, w, h)
}
