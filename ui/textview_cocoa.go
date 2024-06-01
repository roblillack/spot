//go:build !fltk && (darwin || cocoa)

package ui

import (
	"github.com/roblillack/spot"
	"github.com/roblillack/spot/ui/internal/cocoa"
)

type nativeTypeTextView = *cocoa.TextField

func (c *TextView) Update(nextComponent spot.Control) bool {
	next, ok := nextComponent.(*TextView)
	if !ok {
		return false
	}

	if next.Text != c.Text {
		c.Text = next.Text
		c.ref.SetStringValue(c.Text)
	}

	if next.FontSize != c.FontSize && c.FontSize > 0 {
		c.FontSize = next.FontSize
		c.ref.SetFontSize(c.FontSize)
	}

	return true
}

func (c *TextView) Mount(ctx *spot.RenderContext, parent spot.Control) any {
	if c.ref != nil {
		return c.ref
	}

	x, y, w, h := calcLayout(parent, c.X, c.Y, c.Width, c.Height)
	c.ref = cocoa.NewTextField(x, y, w, h)
	c.ref.SetStringValue(c.Text)
	c.ref.SetFontFamily("Arial")
	c.ref.SetEditable(false)
	c.ref.SetSelectable(true)
	if c.FontSize > 0 {
		c.ref.SetFontSize(c.FontSize)
	}

	if window, ok := parent.(*Window); ok && window != nil && window.ref != nil {
		window.ref.AddTextField(c.ref)
	}

	return c.ref
}

func (c *TextView) Unmount() {
	if c.ref == nil {
		return
	}

	c.ref.Remove()
	c.ref = nil
}

func (c *TextView) Layout(ctx *spot.RenderContext, parent spot.Control) {
	x, y, w, h := calcLayout(parent, c.X, c.Y, c.Width, c.Height)
	c.ref.SetFrame(x, y, w, h)
}
