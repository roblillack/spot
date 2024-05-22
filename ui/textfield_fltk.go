//go:build !cocoa && (fltk || !darwin)

package ui

import (
	goFltk "github.com/pwiecz/go-fltk"
	"github.com/roblillack/spot"
)

type nativeTypeTextField = *goFltk.Input

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
			// if w.ref.Buffer() == nil {
			// 	w.ref.SetBuffer(goFltk.NewTextBuffer())
			// }
			// w.ref.Buffer().SetText(w.Value)
			w.ref.SetValue(w.Value)
		}
	}

	// if next.FontSize != w.FontSize {
	// 	w.FontSize = next.FontSize
	// 	if w.ref != nil && w.FontSize > 0 {
	// 		w.ref.SetTextSize(w.FontSize)
	// 	}
	// }

	w.OnChange = next.OnChange

	return true
}

func (w *TextField) Mount(parent spot.Control) any {
	if w.ref != nil {
		return w.ref
	}

	w.ref = goFltk.NewInput(w.X, w.Y, w.Width, w.Height)
	// w.ref.SetBuffer(goFltk.NewTextBuffer())
	// w.ref.Buffer().SetText(w.Value)
	w.ref.SetValue(w.Value)
	// w.ref.Deactivate()
	// w.ref.SetWrapMode(goFltk.WRAP_AT_BOUNDS, 0)
	// if w.FontSize > 0 {
	// 	w.ref.SetTextSize(w.FontSize)
	// }
	w.ref.SetCallback(w.callback)
	w.ref.SetCallbackCondition(goFltk.WhenChanged)

	if window, ok := parent.(*Window); ok && window != nil && window.ref != nil {
		window.ref.Add(w.ref)
	}

	return w.ref
}

func (w *TextField) callback() {
	if w.OnChange != nil {
		val := w.ref.Value()
		if val != w.Value {
			w.Value = val
			w.OnChange(val)
		}
	}
}

var _ spot.Control = &TextField{}
