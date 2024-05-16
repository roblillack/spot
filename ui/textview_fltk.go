//go:build !cocoa && (fltk || !darwin)

package ui

import (
	goFltk "github.com/pwiecz/go-fltk"
	"github.com/roblillack/spot"
)

type TextView struct {
	X      int
	Y      int
	Width  int
	Height int
	Text   string
	ref    *goFltk.TextDisplay
}

var _ spot.HostComponent = &TextView{}

func (w *TextView) Equals(other spot.HostComponent) bool {
	next, ok := other.(*TextView)
	if !ok {
		return false
	}

	return next.Text == w.Text
}

func (w *TextView) Update(nextComponent spot.HostComponent) bool {
	next, ok := nextComponent.(*TextView)
	if !ok {
		return false
	}

	if next.Text != w.Text {
		w.Text = next.Text
		w.ref.SetBuffer(goFltk.NewTextBuffer())
		w.ref.Buffer().SetText(w.Text)
	}

	return true
}

func (w *TextView) Mount(parent spot.HostComponent) any {
	if w.ref != nil {
		return w.ref
	}

	w.ref = goFltk.NewTextDisplay(w.X, w.Y, w.Width, w.Height)
	w.ref.SetBuffer(goFltk.NewTextBuffer())
	w.ref.Buffer().SetText(w.Text)
	w.ref.Deactivate()
	w.ref.SetWrapMode(goFltk.WRAP_AT_BOUNDS, 0)

	if window, ok := parent.(*Window); ok && window != nil && window.ref != nil {
		window.ref.Add(w.ref)
	}

	return w.ref
}
