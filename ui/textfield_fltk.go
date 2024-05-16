//go:build !cocoa && (fltk || !darwin)

package ui

import (
	goFltk "github.com/pwiecz/go-fltk"
	"github.com/roblillack/spot"
)

type TextField struct {
	X        int
	Y        int
	Width    int
	Height   int
	Value    string
	FontSize int
	ref      *goFltk.TextEditor
}

func (w *TextField) Equals(other spot.Control) bool {
	next, ok := other.(*TextField)
	if !ok {
		return false
	}

	if w == nil && next != nil || w != nil && next == nil {
		return false
	}

	return next.Value == w.Value && w.FontSize == next.FontSize
}

func (w *TextField) Update(nextComponent spot.Control) bool {
	next, ok := nextComponent.(*TextField)
	if !ok {
		return false
	}

	if w.ref == nil {
		return false
	}

	if next.Value != w.Value {
		w.Value = next.Value
		if w.ref != nil {
			if w.ref.Buffer() == nil {
				w.ref.SetBuffer(goFltk.NewTextBuffer())
			}
			w.ref.Buffer().SetText(w.Value)
		}
	}

	if next.FontSize != w.FontSize {
		w.FontSize = next.FontSize
		if w.ref != nil {
			w.ref.SetTextSize(w.FontSize)
		}
	}

	return true
}

func (w *TextField) Mount(parent spot.Control) any {
	if w.ref != nil {
		return w.ref
	}

	w.ref = goFltk.NewTextEditor(w.X, w.Y, w.Width, w.Height)
	w.ref.SetBuffer(goFltk.NewTextBuffer())
	w.ref.Buffer().SetText(w.Value)
	// w.ref.Deactivate()
	w.ref.SetWrapMode(goFltk.WRAP_AT_BOUNDS, 0)
	w.ref.SetTextSize(w.FontSize)

	if window, ok := parent.(*Window); ok && window != nil && window.ref != nil {
		window.ref.Add(w.ref)
	}

	return w.ref
}

var _ spot.Control = &TextField{}
