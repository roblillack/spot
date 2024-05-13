//go:build !cocoa && (fltk || !darwin)

package ui

import (
	"journey/spot"

	goFltk "github.com/pwiecz/go-fltk"
)

type Label struct {
	X        int
	Y        int
	Width    int
	Height   int
	Value    string
	FontSize int
	ref      *goFltk.TextDisplay
}

func (w *Label) Equals(other spot.Component) bool {
	next, ok := other.(*Label)
	if !ok {
		return false
	}

	if w == nil && next != nil || w != nil && next == nil {
		return false
	}

	return next.Value == w.Value && w.FontSize == next.FontSize
}

func (w *Label) Update(nextComponent spot.Component) bool {
	next, ok := nextComponent.(*Label)
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

func (w *Label) Mount() any {
	if w.ref != nil {
		return w.ref
	}

	w.ref = goFltk.NewTextDisplay(w.X, w.Y, w.Width, w.Height)
	buf := goFltk.NewTextBuffer()
	buf.SetText(w.Value)
	w.ref.SetBuffer(buf)
	w.ref.SetTextSize(w.FontSize)
	w.ref.HideCursor()
	return w.ref
}
