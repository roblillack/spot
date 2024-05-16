//go:build !cocoa && (fltk || !darwin)

package ui

import (
	goFltk "github.com/pwiecz/go-fltk"
	"github.com/roblillack/spot"
)

type nativeTypeTextView = *goFltk.TextDisplay

func (w *TextView) Update(nextComponent spot.Control) bool {
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

func (w *TextView) Mount(parent spot.Control) any {
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
