//go:build !cocoa && (fltk || !darwin)

package ui

import (
	goFltk "github.com/pwiecz/go-fltk"
	"github.com/roblillack/spot"
)

type nativeTypeTextView = *goFltk.TextDisplay

func (c *TextView) Update(nextComponent spot.Control) bool {
	next, ok := nextComponent.(*TextView)
	if !ok {
		return false
	}

	if next.Text != c.Text {
		c.Text = next.Text
		c.ref.SetBuffer(goFltk.NewTextBuffer())
		c.ref.Buffer().SetText(c.Text)
	}

	return true
}

func (c *TextView) Mount(ctx *spot.RenderContext, parent spot.Control) any {
	if c.ref != nil {
		return c.ref
	}

	x, y, w, h := calcLayout(parent, c.X, c.Y, c.Width, c.Height)
	c.ref = goFltk.NewTextDisplay(x, y, w, h)
	c.ref.SetBuffer(goFltk.NewTextBuffer())
	c.ref.Buffer().SetText(c.Text)
	// c.ref.Deactivate()
	c.ref.SetWrapMode(goFltk.WRAP_AT_BOUNDS, 0)

	if window, ok := parent.(*Window); ok && window != nil && window.ref != nil {
		window.ref.Add(c.ref)
	}

	return c.ref
}

func (c *TextView) Unmount() {
	if c.ref == nil {
		return
	}

	c.ref.Destroy()
	c.ref = nil
}

func (c *TextView) Layout(ctx *spot.RenderContext, parent spot.Control) {
	x, y, w, h := calcLayout(parent, c.X, c.Y, c.Width, c.Height)
	c.ref.Resize(x, y, w, h)
}
