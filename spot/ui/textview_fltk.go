//go:build !cocoa && (fltk || !darwin)

package ui

import (
	"journey/spot"

	goFltk "github.com/pwiecz/go-fltk"
)

type TextView struct {
	X      int
	Y      int
	Width  int
	Height int
	Text   string
	ref    *goFltk.TextDisplay
}

var _ spot.Component = &TextView{}

func (w *TextView) Equals(other spot.Component) bool {
	next, ok := other.(*TextView)
	if !ok {
		return false
	}

	return next.Text == w.Text
}

func (w *TextView) Update(nextComponent spot.Component) bool {
	next, ok := nextComponent.(*TextView)
	if !ok {
		return false
	}

	if next.Text != w.Text {
		w.Text = next.Text
		w.ref.SetLabel(w.Text)
	}

	return true
}

func (w *TextView) Mount() any {
	if w.ref != nil {
		return w.ref
	}

	w.ref = goFltk.NewTextDisplay(w.X, w.Y, w.Width, w.Height, w.Text)
	w.ref.Deactivate()
	w.ref.SetWrapMode(goFltk.WRAP_AT_BOUNDS, 0)
	return w.ref
}
