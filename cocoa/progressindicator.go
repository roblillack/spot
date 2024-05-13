package cocoa

import (
	"fmt"
	"journey/spot"
	"time"

	"github.com/mojbro/gocoa"
)

type ProgressIndicator struct {
	X             int
	Y             int
	Width         int
	Height        int
	Min           float64
	Max           float64
	Value         float64
	Indeterminate bool
	ref           *gocoa.ProgressIndicator
}

var _ spot.Component = &ProgressIndicator{}

func (w *ProgressIndicator) Mount() any {
	if w.ref != nil {
		return w.ref
	}

	startTime := time.Now()
	w.ref = gocoa.NewProgressIndicator(w.X, w.Y, w.Width, w.Height)
	fmt.Printf("ProgressIndicator: %s\n", time.Since(startTime))
	w.ref.SetLimits(w.Min, w.Max)
	w.ref.SetValue(w.Value)
	w.ref.SetIsIndeterminate(w.Indeterminate)
	// w.ref.Show()
	return w.ref
}

func (w *ProgressIndicator) Equals(other spot.Component) bool {
	next, ok := other.(*ProgressIndicator)
	if !ok {
		return false
	}

	if w == nil && next != nil || w != nil && next == nil {
		return false
	}

	return next.Width == w.Width && next.Height == w.Height &&
		next.Min == w.Min && next.Max == w.Max &&
		next.Value == w.Value && next.Indeterminate == w.Indeterminate
}

func (w *ProgressIndicator) Update(nextComponent spot.Component) bool {
	next, ok := nextComponent.(*ProgressIndicator)
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
