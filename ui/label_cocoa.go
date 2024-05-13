//go:build !fltk && (darwin || cocoa)

package ui

import (
	"github.com/mojbro/gocoa"
	"github.com/roblillack/spot"
)

type Label struct {
	X        int
	Y        int
	Width    int
	Height   int
	Value    string
	FontSize int
	ref      *gocoa.TextView
}

func (w *Label) Equals(other spot.Component) bool {
	next, ok := other.(*Label)
	if !ok {
		return false
	}

	return next.Value == w.Value
}

func (w *Label) Update(nextComponent spot.Component) bool {
	next, ok := nextComponent.(*Label)
	if !ok {
		return false
	}

	if next.Value != w.Value {
		w.Value = next.Value
		w.ref.SetText(w.Value)
	}

	return true
}

func (w *Label) Mount() any {
	if w.ref != nil {
		return w.ref
	}

	w.ref = gocoa.NewTextView(w.X, w.Y, w.Width, w.Height)
	w.ref.SetText(w.Value)
	return w.ref
}
