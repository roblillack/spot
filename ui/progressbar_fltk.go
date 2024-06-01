//go:build !cocoa && (fltk || !darwin)

package ui

import (
	goFltk "github.com/pwiecz/go-fltk"
	"github.com/roblillack/spot"
)

type nativeTypeProgressBar = *goFltk.Progress

func (b *ProgressBar) Update(nextComponent spot.Control) bool {
	next, ok := nextComponent.(*ProgressBar)
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

	return true
}

func (c *ProgressBar) Mount(ctx *spot.RenderContext, parent spot.Control) any {
	if c.ref != nil {
		return c.ref
	}

	x, y, w, h := calcLayout(parent, c.X, c.Y, c.Width, c.Height)
	c.ref = goFltk.NewProgress(x, y, w, h)
	c.ref.SetMaximum(c.Max)
	c.ref.SetMinimum(c.Min)
	c.ref.SetValue(c.Value)
	// b.ref.SetCallback(func() {
	// 	if b.OnValueChanged != nil {
	// 		b.OnValueChanged(b.ref.Value())
	// 	}
	// })

	if window, ok := parent.(*Window); ok && window != nil && window.ref != nil {
		window.ref.Add(c.ref)
	}

	return c.ref
}

func (c *ProgressBar) Unmount() {
	if c.ref == nil {
		return
	}

	c.ref.Destroy()
	c.ref = nil
}

func (c *ProgressBar) Layout(ctx *spot.RenderContext, parent spot.Control) {
	if c.ref == nil {
		return
	}

	x, y, w, h := calcLayout(parent, c.X, c.Y, c.Width, c.Height)
	c.ref.Resize(x, y, w, h)
}
