//go:build !cocoa && (fltk || !darwin)

package ui

import (
	goFltk "github.com/pwiecz/go-fltk"
	"github.com/roblillack/spot"
)

type nativeTypeLabel = *goFltk.Box

func (w *Label) Update(nextComponent spot.Control) bool {
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
			// if w.ref.Buffer() == nil {
			// 	w.ref.SetBuffer(goFltk.NewTextBuffer())
			// }
			// w.ref.Buffer().SetText(w.Value)
			w.ref.SetLabel(w.Value)
		}
	}

	if next.FontSize != w.FontSize {
		w.FontSize = next.FontSize
		if w.ref != nil && w.FontSize > 0 {
			w.ref.SetLabelSize(w.FontSize)
		}
	}

	if next.Align != w.Align {
		w.setAlign(next.Align)
	}

	return true
}

func (w *Label) setAlign(a LabelAlignment) {
	w.Align = a
	if w.ref == nil {
		return
	}
	switch a {
	case LabelAlignmentLeft:
		w.ref.SetAlign(goFltk.ALIGN_INSIDE | goFltk.ALIGN_LEFT)
	case LabelAlignmentCenter:
		w.ref.SetAlign(goFltk.ALIGN_INSIDE | goFltk.ALIGN_CENTER)
	case LabelAlignmentRight:
		w.ref.SetAlign(goFltk.ALIGN_INSIDE | goFltk.ALIGN_RIGHT)
	}
}

func (w *Label) Mount(parent spot.Control) any {
	if w.ref != nil {
		return w.ref
	}

	// w.ref = goFltk.NewTextDisplay(w.X, w.Y, w.Width, w.Height)
	w.ref = goFltk.NewBox(goFltk.NO_BOX, w.X, w.Y, w.Width, w.Height)
	w.ref.SetLabel(w.Value)
	w.setAlign(w.Align)
	// buf := goFltk.NewTextBuffer()
	// buf.SetText(w.Value)
	// w.ref.SetBuffer(buf)
	if w.FontSize > 0 {
		// w.ref.SetTextSize(w.FontSize)
		w.ref.SetLabelSize(w.FontSize)
	}
	// w.ref.HideCursor()

	if window, ok := parent.(*Window); ok && window != nil && window.ref != nil {
		window.ref.Add(w.ref)
	}

	return w.ref
}

func (w *Label) Unmount() {
	if w.ref != nil {
		w.ref.Destroy()
		w.ref = nil
	}
}
