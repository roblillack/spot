//go:build !fltk && (darwin || cocoa)

package ui

import (
	"github.com/roblillack/gocoa"
	"github.com/roblillack/spot"
)

type nativeTypeLabel = *gocoa.TextField

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
		w.ref.SetStringValue(w.Value)
	}

	if next.FontSize != w.FontSize {
		w.FontSize = next.FontSize
		if w.FontSize > 0 {
			w.ref.SetFontSize(w.FontSize)
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
		w.ref.SetAlignmentLeft()
	case LabelAlignmentCenter:
		w.ref.SetAlignmentCenter()
	case LabelAlignmentRight:
		w.ref.SetAlignmentRight()
	}
}

func (w *Label) Mount(parent spot.Control) any {
	if w.ref != nil {
		return w.ref
	}

	w.ref = gocoa.NewTextField(w.X, w.Y, w.Width, w.Height)
	w.ref.SetBezeled(false)
	w.ref.SetDrawsBackground(false)
	w.ref.SetEditable(false)
	w.ref.SetSelectable(false)
	w.ref.SetStringValue(w.Value)
	w.setAlign(w.Align)
	if w.FontSize > 0 {
		w.ref.SetFontSize(w.FontSize)
	}

	if window, ok := parent.(*Window); ok && window != nil && window.ref != nil {
		window.ref.AddTextField(w.ref)
	}

	return w.ref
}

func (c *Label) Unmount() {
	if c.ref == nil {
		return
	}

	c.ref.Remove()
	c.ref = nil
}
