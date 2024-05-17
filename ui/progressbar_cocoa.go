//go:build !fltk && (darwin || cocoa)

package ui

import (
	"github.com/roblillack/gocoa"
	"github.com/roblillack/spot"
)

type nativeTypeProgressBar = *gocoa.ProgressIndicator

func (w *ProgressBar) Mount(parent spot.Control) any {
	if w.ref != nil {
		return w.ref
	}

	w.ref = gocoa.NewProgressIndicator(w.X, w.Y, w.Width, w.Height)
	w.ref.SetLimits(w.Min, w.Max)
	w.ref.SetValue(w.Value)
	w.ref.SetIsIndeterminate(w.Indeterminate)
	// w.ref.Show()

	if window, ok := parent.(*Window); ok && window != nil && window.ref != nil {
		window.ref.AddProgressIndicator(w.ref)
	}

	return w.ref
}

func (w *ProgressBar) Update(nextComponent spot.Control) bool {
	next, ok := nextComponent.(*ProgressBar)
	if !ok {
		return false
	}

	if w.ref == nil {
		return false
	}

	if next.Max != w.Max || next.Min != w.Min {
		w.Min = next.Min
		w.Max = next.Max
		w.ref.SetLimits(w.Min, w.Max)
	}

	if next.Value != w.Value {
		w.Value = next.Value
		w.ref.SetValue(w.Value)
	}

	if next.Indeterminate != w.Indeterminate {
		w.Indeterminate = next.Indeterminate
		w.ref.SetIsIndeterminate(w.Indeterminate)
	}

	return true
}
