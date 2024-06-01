//go:build !fltk && (darwin || cocoa)

package ui

import (
	"github.com/roblillack/spot"
	"github.com/roblillack/spot/ui/internal/cocoa"
)

type nativeTypeLabel = *cocoa.TextField

func (w *Label) Update(nextComponent spot.Control) bool {
	next, ok := nextComponent.(*Label)
	if !ok {
		return false
	}

	if w.ref == nil {
		return false
	}

	if next.Value != w.Value {
		w.Value = next.Value
		w.ref.SetStringValue(w.Value)
	}

	if next.FontSize != w.FontSize {
		w.FontSize = next.FontSize
		if w.FontSize > 0 {
			w.ref.SetFontSize(w.FontSize)
		}
	}

	if next.Align != w.Align {
		w.setAlign(next.Align)
	}

	return true
}

func (w *Label) setAlign(a LabelAlignment) {
	w.Align = a
	if w.ref == nil {
		return
	}

	switch a {
	case LabelAlignmentLeft:
		w.ref.SetAlignmentLeft()
	case LabelAlignmentCenter:
		w.ref.SetAlignmentCenter()
	case LabelAlignmentRight:
		w.ref.SetAlignmentRight()
	}
}

func (c *Label) Mount(ctx *spot.RenderContext, parent spot.Control) any {
	if c.ref != nil {
		return c.ref
	}

	x, y, w, h := calcLayout(parent, c.X, c.Y, c.Width, c.Height)
	c.ref = cocoa.NewTextField(x, y, w, h)
	c.ref.SetBezeled(false)
	c.ref.SetDrawsBackground(true)
	c.ref.SetEditable(false)
	c.ref.SetSelectable(false)
	c.ref.SetStringValue(c.Value)
	c.setAlign(c.Align)
	if c.FontSize > 0 {
		c.ref.SetFontSize(c.FontSize)
	}

	if window, ok := parent.(*Window); ok && window != nil && window.ref != nil {
		window.ref.AddTextField(c.ref)
	}

	return c.ref
}

func (c *Label) Unmount() {
	if c.ref == nil {
		return
	}

	c.ref.Remove()
	c.ref = nil
}

func (c *Label) Layout(ctx *spot.RenderContext, parent spot.Control) {
	x, y, w, h := calcLayout(parent, c.X, c.Y, c.Width, c.Height)
	c.ref.SetFrame(x, y, w, h)
}
