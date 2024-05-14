//go:build !fltk && (darwin || cocoa)

package ui

import (
	"github.com/roblillack/gocoa"
	"github.com/roblillack/spot"
)

type nativeTypeLabel = *gocoa.TextView

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
		w.ref.SetText(w.Value)
	}

	if next.FontSize != w.FontSize {
		w.FontSize = next.FontSize
		if w.FontSize > 0 {
			w.ref.SetFontSize(w.FontSize)
		}
	}

	return true
}

func (w *Label) Mount() any {
	if w.ref != nil {
		return w.ref
	}

	w.ref = gocoa.NewTextView(w.X, w.Y, w.Width, w.Height)
	w.ref.SetText(w.Value)
	if w.FontSize > 0 {
		w.ref.SetFontSize(w.FontSize)
	}
	return w.ref
}
