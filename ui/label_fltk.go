//go:build !cocoa && (fltk || !darwin)

package ui

import (
	goFltk "github.com/pwiecz/go-fltk"
	"github.com/roblillack/spot"
)

type nativeTypeLabel = *goFltk.TextDisplay

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
		if w.ref != nil && w.FontSize > 0 {
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
