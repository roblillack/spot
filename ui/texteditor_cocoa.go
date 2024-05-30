//go:build !fltk && (darwin || cocoa)

package ui

import (
	"github.com/roblillack/gocoa"
	"github.com/roblillack/spot"
)

type nativeTypeTextEditor = *gocoa.TextField

func (w *TextEditor) Update(nextComponent spot.Control) bool {
	next, ok := nextComponent.(*TextEditor)
	if !ok {
		return false
	}

	if next.Text != w.Text {
		w.Text = next.Text
		w.ref.SetStringValue(w.Text)
	}

	if next.FontSize != w.FontSize && w.FontSize > 0 {
		w.FontSize = next.FontSize
		w.ref.SetFontSize(w.FontSize)
	}

	return true
}

func (w *TextEditor) Mount(parent spot.Control) any {
	if w.ref != nil {
		return w.ref
	}

	w.ref = gocoa.NewTextField(w.X, w.Y, w.Width, w.Height)
	w.ref.SetStringValue(w.Text)
	w.ref.SetFontFamily("Arial")
	w.ref.SetEditable(false)
	w.ref.SetSelectable(true)
	if w.FontSize > 0 {
		w.ref.SetFontSize(w.FontSize)
	}

	if window, ok := parent.(*Window); ok && window != nil && window.ref != nil {
		window.ref.AddTextField(w.ref)
	}

	return w.ref
}
