//go:build !cocoa && (fltk || !darwin)

package ui

import (
	goFltk "github.com/pwiecz/go-fltk"
	"github.com/roblillack/spot"
)

type nativeTypeSpinner = *goFltk.Spinner

func (b *Spinner) Update(nextComponent spot.Mountable) bool {
	next, ok := nextComponent.(*Spinner)
	if !ok {
		return false
	}

	if b.ref == nil {
		return false
	}

	if next.Min != b.Min {
		b.Min = next.Min
		b.ref.SetMinimum(b.Min)
	}

	if next.Max != b.Max {
		b.Max = next.Max
		b.ref.SetMaximum(b.Max)
	}

	if next.Value != b.Value {
		b.Value = next.Value
		b.ref.SetValue(b.Value)
	}

	if next.Step != b.Step {
		b.Step = next.Step
		b.ref.SetStep(b.Step)
	}

	return true
}

func (c *Spinner) Mount(ctx *spot.RenderContext, parent spot.Mountable) any {
	if c.ref != nil {
		return c.ref
	}

	x, y, w, h := CalcLayout(parent, c.X, c.Y, c.Width, c.Height)
	c.ref = goFltk.NewSpinner(x, y, w, h)
	c.ref.SetMaximum(c.Max)
	c.ref.SetMinimum(c.Min)
	c.ref.SetValue(c.Value)
	c.ref.SetStep(c.Step)
	c.ref.SetCallback(func() {
		if c.OnValueChanged != nil {
			c.OnValueChanged(c.ref.Value())
		}
	})

	if window, ok := parent.(*Window); ok && window != nil && window.ref != nil {
		window.ref.Add(c.ref)
	}

	return c.ref
}

func (c *Spinner) Unmount() {
	if c.ref == nil {
		return
	}

	c.ref.Destroy()
	c.ref = nil
}

func (c *Spinner) Layout(ctx *spot.RenderContext, parent spot.Container) {
	if c.ref == nil {
		return
	}

	x, y, w, h := CalcLayout(parent, c.X, c.Y, c.Width, c.Height)
	c.ref.Resize(x, y, w, h)
}
