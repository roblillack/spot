package cocoa

import (
	"journey/spot"

	"github.com/mojbro/gocoa"
)

type TextField struct {
	X            int
	Y            int
	Width        int
	Height       int
	Editable     bool
	Bezeled      bool
	Selectable   bool
	Value        string
	FontSize     int
	NoBackground bool
	ref          *gocoa.TextField
}

var _ spot.Component = &TextField{}

func (w *TextField) Equals(other spot.Component) bool {
	next, ok := other.(*TextField)
	if !ok {
		return false
	}

	return next.Value == w.Value &&
		next.Editable == w.Editable &&
		next.Bezeled == w.Bezeled &&
		next.Selectable == w.Selectable &&
		next.FontSize == w.FontSize &&
		next.NoBackground == w.NoBackground

}

func (w *TextField) Update(nextComponent spot.Component) bool {
	next, ok := nextComponent.(*TextField)
	if !ok {
		return false
	}

	if next.Value != w.Value {
		w.Value = next.Value
		if w.ref != nil {
			w.ref.SetStringValue(w.Value)
		}
	}

	if next.Editable != w.Editable {
		w.Editable = next.Editable
		if w.ref != nil {
			w.ref.SetEditable(w.Editable)
		}
	}

	if next.Bezeled != w.Bezeled {
		w.Bezeled = next.Bezeled
		if w.ref != nil {
			w.ref.SetBezeled(w.Bezeled)
		}
	}

	if next.Selectable != w.Selectable {
		w.Selectable = next.Selectable
		if w.ref != nil {
			w.ref.SetSelectable(w.Selectable)
		}
	}

	if next.FontSize != w.FontSize {
		w.FontSize = next.FontSize
		if w.ref != nil {
			w.ref.SetFontSize(w.FontSize)
		}
	}

	if next.NoBackground != w.NoBackground {
		w.NoBackground = next.NoBackground
		if w.ref != nil {
			w.ref.SetDrawsBackground(!w.NoBackground)
		}
	}

	return true
}

func (w *TextField) Mount() any {
	if w.ref != nil {
		return w.ref
	}

	w.ref = gocoa.NewTextField(w.X, w.Y, w.Width, w.Height)
	w.ref.SetEditable(w.Editable)
	w.ref.SetBezeled(w.Bezeled)
	w.ref.SetSelectable(w.Selectable)
	w.ref.SetStringValue(w.Value)
	w.ref.SetFontSize(w.FontSize)
	w.ref.SetDrawsBackground(!w.NoBackground)
	return w.ref
}
