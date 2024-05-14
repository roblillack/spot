//go:build !fltk && (darwin || cocoa)

package ui

import (
	"github.com/roblillack/gocoa"
	"github.com/roblillack/spot"
)

type TextView struct {
	X      int
	Y      int
	Width  int
	Height int
	Text   string
	ref    *gocoa.TextView
}

var _ spot.Component = &Label{}

func (w *TextView) Equals(other spot.Component) bool {
	next, ok := other.(*Label)
	if !ok {
		return false
	}

	return next.Value == w.Text
}

func (w *TextView) Update(nextComponent spot.Component) bool {
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

func (w *TextView) Mount() any {
	if w.ref != nil {
		return w.ref
	}

	w.ref = gocoa.NewTextView(w.X, w.Y, w.Width, w.Height)
	w.ref.SetText(w.Text)
	return w.ref
}
