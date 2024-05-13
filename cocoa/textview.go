package cocoa

import (
	"journey/spot"

	"github.com/mojbro/gocoa"
)

type TextView struct {
	X      int
	Y      int
	Width  int
	Height int
	Text   string
	ref    *gocoa.TextView
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
