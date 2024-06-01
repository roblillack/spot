//go:build !fltk && (darwin || cocoa)

package ui

import (
	"github.com/roblillack/spot"
	"github.com/roblillack/spot/ui/internal/cocoa"
)

type nativeTypeTextEditor = *cocoa.TextView

func (w *TextEditor) Update(nextComponent spot.Control) bool {
	next, ok := nextComponent.(*TextEditor)
	if !ok {
		return false
	}

	if next.Text != w.Text {
		w.Text = next.Text
		w.ref.SetText(w.Text)
	}

	if next.FontSize != w.FontSize && w.FontSize > 0 {
		w.FontSize = next.FontSize
		w.ref.SetFontSize(w.FontSize)
	}

	return true
}

func (w *TextEditor) Mount(ctx *spot.RenderContext, parent spot.Control) any {
	if w.ref != nil {
		return w.ref
	}

	w.ref = cocoa.NewTextView(w.X, w.Y, w.Width, w.Height)
	w.ref.SetText(w.Text)
	// w.ref.SetFontFamily("Arial")
	w.ref.SetEditable(true)
	// w.ref.SetSelectable(true)
	if w.FontSize > 0 {
		w.ref.SetFontSize(w.FontSize)
	}

	if window, ok := parent.(*Window); ok && window != nil && window.ref != nil {
		window.ref.AddTextView(w.ref)
	}

	return w.ref
}

func (c *TextEditor) Unmount() {
	if c.ref == nil {
		return
	}

	c.ref.Remove()
	c.ref = nil
}

func (c *TextEditor) Layout(ctx *spot.RenderContext, parent spot.Control) {
	x, y, w, h := calcLayout(parent, c.X, c.Y, c.Width, c.Height)
	c.ref.SetFrame(x, y, w, h)
}
