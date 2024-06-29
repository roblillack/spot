//go:build !cocoa && (fltk || !darwin)

package ui

import (
	goFltk "github.com/pwiecz/go-fltk"
	"github.com/roblillack/spot"
)

type nativeTypeTextField = *goFltk.Input

func (c *TextField) Update(nextComponent spot.Mountable) bool {
	next, ok := nextComponent.(*TextField)
	if !ok {
		return false
	}

	if c.ref == nil {
		return false
	}

	if next.Value != c.Value {
		c.Value = next.Value
		if c.ref != nil {
			// if w.ref.Buffer() == nil {
			// 	w.ref.SetBuffer(goFltk.NewTextBuffer())
			// }
			// w.ref.Buffer().SetText(w.Value)
			c.ref.SetValue(c.Value)
		}
	}

	// if next.FontSize != w.FontSize {
	// 	w.FontSize = next.FontSize
	// 	if w.ref != nil && w.FontSize > 0 {
	// 		w.ref.SetTextSize(w.FontSize)
	// 	}
	// }

	c.OnChange = next.OnChange

	return true
}

func (c *TextField) Mount(ctx *spot.RenderContext, parent spot.Mountable) any {
	if c.ref != nil {
		return c.ref
	}

	x, y, w, h := CalcLayout(parent, c.X, c.Y, c.Width, c.Height)
	c.ref = goFltk.NewInput(x, y, w, h)
	// w.ref.SetBuffer(goFltk.NewTextBuffer())
	// w.ref.Buffer().SetText(w.Value)
	c.ref.SetValue(c.Value)
	// w.ref.Deactivate()
	// w.ref.SetWrapMode(goFltk.WRAP_AT_BOUNDS, 0)
	// if w.FontSize > 0 {
	// 	w.ref.SetTextSize(w.FontSize)
	// }
	c.ref.SetCallback(c.callback)
	c.ref.SetCallbackCondition(goFltk.WhenChanged)

	if window, ok := parent.(*Window); ok && window != nil && window.ref != nil {
		window.ref.Add(c.ref)
	}

	return c.ref
}

func (c *TextField) callback() {
	if c.OnChange != nil {
		val := c.ref.Value()
		if val != c.Value {
			c.Value = val
			c.OnChange(val)
		}
	}
}

func (c *TextField) Unmount() {
	if c.ref == nil {
		return
	}

	c.ref.Destroy()
	c.ref = nil
}

func (c *TextField) Layout(ctx *spot.RenderContext, parent spot.Container) {
	x, y, w, h := CalcLayout(parent, c.X, c.Y, c.Width, c.Height)
	c.ref.Resize(x, y, w, h)
}
