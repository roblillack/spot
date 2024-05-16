//go:build !fltk && (darwin || cocoa)

package ui

import (
	"github.com/roblillack/gocoa"
	"github.com/roblillack/spot"
)

type nativeTypeTextField = *gocoa.TextField

func (w *TextField) Update(nextComponent spot.Control) bool {
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

	// if next.Editable != w.Editable {
	// 	w.Editable = next.Editable
	// 	if w.ref != nil {
	// 		w.ref.SetEditable(w.Editable)
	// 	}
	// }

	// if next.Bezeled != w.Bezeled {
	// 	w.Bezeled = next.Bezeled
	// 	if w.ref != nil {
	// 		w.ref.SetBezeled(w.Bezeled)
	// 	}
	// }

	// if next.Selectable != w.Selectable {
	// 	w.Selectable = next.Selectable
	// 	if w.ref != nil {
	// 		w.ref.SetSelectable(w.Selectable)
	// 	}
	// }

	if next.FontSize != w.FontSize {
		w.FontSize = next.FontSize
		if w.ref != nil {
			w.ref.SetFontSize(w.FontSize)
		}
	}

	// if next.NoBackground != w.NoBackground {
	// 	w.NoBackground = next.NoBackground
	// 	if w.ref != nil {
	// 		w.ref.SetDrawsBackground(!w.NoBackground)
	// 	}
	// }

	return true
}

func (w *TextField) Mount(parent spot.Control) any {
	if w.ref != nil {
		return w.ref
	}

	w.ref = gocoa.NewTextField(w.X, w.Y, w.Width, w.Height)
	// w.ref.SetEditable(w.Editable)
	// w.ref.SetBezeled(w.Bezeled)
	// w.ref.SetSelectable(w.Selectable)
	w.ref.SetStringValue(w.Value)
	w.ref.SetFontFamily("Arial")
	w.ref.SetFontSize(w.FontSize)
	// w.ref.SetDrawsBackground(!w.NoBackground)

	if window, ok := parent.(*Window); ok && window != nil && window.ref != nil {
		window.ref.AddTextField(w.ref)
	}

	return w.ref
}
