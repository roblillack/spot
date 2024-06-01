//go:build !fltk && (darwin || cocoa)

package ui

import (
	"github.com/roblillack/spot"
	"github.com/roblillack/spot/ui/internal/cocoa"
)

type nativeTypeSlider = *cocoa.Slider

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

func (c *Slider) Mount(ctx *spot.RenderContext, parent spot.Control) any {
	if c.ref != nil {
		return c.ref
	}

	x, y, w, h := calcLayout(parent, c.X, c.Y, c.Width, c.Height)
	c.ref = cocoa.NewSlider(x, y, w, h)
	c.ref.SetMaximumValue(c.Max)
	c.ref.SetMinimumValue(c.Min)
	c.ref.SetValue(c.Value)
	c.ref.OnSliderValueChanged(c.callback)

	if window, ok := parent.(*Window); ok && window != nil && window.ref != nil {
		window.ref.AddSlider(c.ref)
	}

	return c.ref
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

func (c *Slider) Layout(ctx *spot.RenderContext, parent spot.Control) {
	x, y, w, h := calcLayout(parent, c.X, c.Y, c.Width, c.Height)
	c.ref.SetFrame(x, y, w, h)
}
