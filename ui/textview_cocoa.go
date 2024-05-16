//go:build !fltk && (darwin || cocoa)

package ui

import (
	"github.com/roblillack/gocoa"
	"github.com/roblillack/spot"
)

type nativeTypeTextView = *gocoa.TextView

func (w *TextView) Update(nextComponent spot.Control) bool {
	next, ok := nextComponent.(*TextView)
	if !ok {
		return false
	}

	if next.Text != w.Text {
		w.Text = next.Text
		w.ref.SetText(w.Text)
	}

	return true
}

func (w *TextView) Mount(parent spot.Control) any {
	if w.ref != nil {
		return w.ref
	}

	w.ref = gocoa.NewTextView(w.X, w.Y, w.Width, w.Height)
	w.ref.SetText(w.Text)

	if window, ok := parent.(*Window); ok && window != nil && window.ref != nil {
		window.ref.AddTextView(w.ref)
	}

	return w.ref
}
