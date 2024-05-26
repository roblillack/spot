//go:build !cocoa && (fltk || !darwin)

package ui

import (
	"log"

	goFltk "github.com/pwiecz/go-fltk"
	"github.com/roblillack/spot"
)

type nativeTypeTextEditor = *goFltk.TextEditor

func (w *TextEditor) Update(nextComponent spot.Control) bool {
	next, ok := nextComponent.(*TextEditor)
	if !ok {
		return false
	}

	log.Println("Update check")
	if next.Text != w.Text {
		log.Println("Updating ...")
		w.Text = next.Text
		w.ref.SetBuffer(goFltk.NewTextBuffer())
		w.ref.Buffer().SetText(w.Text)
	}

	w.OnChange = next.OnChange

	return true
}

func (w *TextEditor) Mount(parent spot.Control) any {
	if w.ref != nil {
		return w.ref
	}

	w.ref = goFltk.NewTextEditor(w.X, w.Y, w.Width, w.Height)
	w.ref.SetBuffer(goFltk.NewTextBuffer())
	w.ref.Buffer().SetText(w.Text)
	// w.ref.Deactivate()
	w.ref.SetWrapMode(goFltk.WRAP_AT_BOUNDS, 0)
	w.ref.SetCallback(w.callback)
	w.ref.SetCallbackCondition(goFltk.WhenChanged)

	if window, ok := parent.(*Window); ok && window != nil && window.ref != nil {
		window.ref.Add(w.ref)
	}

	return w.ref
}

func (w *TextEditor) callback() {
	if w.OnChange != nil {
		val := w.ref.Buffer().Text()
		if val != w.Text {
			w.Text = val
			w.OnChange(val)
		}
	}
}
