//go:build !fltk && (darwin || cocoa)

package ui

import (
	"fmt"

	"github.com/roblillack/gocoa"
	"github.com/roblillack/spot"
)

type Spinner struct {
	X              int
	Y              int
	Width          int
	Height         int
	Min            float64
	Max            float64
	Step           float64
	Value          float64
	OnValueChanged func(float64)
	ref            *gocoa.TextField
}

func (w *Spinner) Update(nextComponent spot.Component) bool {
	next, ok := nextComponent.(*Spinner)
	if !ok {
		return false
	}

	if next.Value != w.Value {
		w.Value = next.Value
		if w.ref != nil {
			w.ref.SetStringValue(fmt.Sprintf("%f", w.Value))
		}
	}

	if next.Min != w.Min {
		w.Min = next.Min
	}

	if next.Max != w.Max {
		w.Max = next.Max
	}

	if next.Step != w.Step {
		w.Step = next.Step
	}

	return true
}

func (w *Spinner) Mount() any {
	if w.ref != nil {
		return w.ref
	}

	w.ref = gocoa.NewTextField(w.X, w.Y, w.Width, w.Height)
	w.ref.SetStringValue(fmt.Sprintf("%f", w.Value))

	return w.ref
}
