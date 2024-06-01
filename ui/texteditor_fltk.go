//go:build !cocoa && (fltk || !darwin)

package ui

import (
	goFltk "github.com/pwiecz/go-fltk"
	"github.com/roblillack/spot"
)

type nativeTypeTextEditor = *goFltk.TextEditor

func (c *TextEditor) Update(nextComponent spot.Control) bool {
	next, ok := nextComponent.(*TextEditor)
	if !ok {
		return false
	}

	if next.Text != c.Text {
		c.Text = next.Text
		c.ref.SetBuffer(goFltk.NewTextBuffer())
		c.ref.Buffer().SetText(c.Text)
	}

	c.OnChange = next.OnChange

	return true
}

func (c *TextEditor) Mount(ctx *spot.RenderContext, parent spot.Control) any {
	if c.ref != nil {
		return c.ref
	}

	x, y, w, h := calcLayout(parent, c.X, c.Y, c.Width, c.Height)
	c.ref = goFltk.NewTextEditor(x, y, w, h)
	c.ref.SetBuffer(goFltk.NewTextBuffer())
	c.ref.Buffer().SetText(c.Text)
	// w.ref.Deactivate()
	c.ref.SetWrapMode(goFltk.WRAP_AT_BOUNDS, 0)
	c.ref.SetCallback(c.callback)
	c.ref.SetCallbackCondition(goFltk.WhenChanged)

	if window, ok := parent.(*Window); ok && window != nil && window.ref != nil {
		window.ref.Add(c.ref)
	}

	return c.ref
}

func (c *TextEditor) callback() {
	if c.OnChange != nil {
		val := c.ref.Buffer().Text()
		if val != c.Text {
			c.Text = val
			c.OnChange(val)
		}
	}
}

func (c *TextEditor) Unmount() {
	if c.ref == nil {
		return
	}

	c.ref.Destroy()
	c.ref = nil
}

func (c *TextEditor) Layout(ctx *spot.RenderContext, parent spot.Control) {
	x, y, w, h := calcLayout(parent, c.X, c.Y, c.Width, c.Height)
	c.ref.Resize(x, y, w, h)
}
