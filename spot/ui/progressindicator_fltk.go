//go:build !cocoa && (fltk || !darwin)

package ui

import (
	"journey/spot"

	goFltk "github.com/pwiecz/go-fltk"
)

type ProgressIndicator struct {
	X              int
	Y              int
	Width          int
	Height         int
	Min            float64
	Max            float64
	Value          float64
	OnValueChanged func(float64)
	ref            *goFltk.Progress
}

func (b *ProgressIndicator) Equals(other spot.Component) bool {
	next, ok := other.(*ProgressIndicator)
	if !ok {
		return false
	}

	if b == nil && next != nil || b != nil && next == nil {
		return false
	}

	return next.Max == b.Max && next.Min == b.Min &&
		next.Value == b.Value
}

func (b *ProgressIndicator) Update(nextComponent spot.Component) bool {
	next, ok := nextComponent.(*ProgressIndicator)
	if !ok {
		return false
	}

	if b.ref == nil {
		return false
	}

	if next.Min != b.Min {
		b.Min = next.Min
		b.ref.SetMinimum(b.Min)
	}

	if next.Max != b.Max {
		b.Max = next.Max
		b.ref.SetMaximum(b.Max)
	}

	if next.Value != b.Value {
		b.Value = next.Value
		b.ref.SetValue(b.Value)
	}

	return true
}

func (b *ProgressIndicator) Mount() any {
	if b.ref != nil {
		return b.ref
	}

	b.ref = goFltk.NewProgress(b.X, b.Y, b.Width, b.Height)
	b.ref.SetMaximum(b.Max)
	b.ref.SetMinimum(b.Min)
	b.ref.SetValue(b.Value)
	b.ref.SetCallback(func() {
		if b.OnValueChanged != nil {
			b.OnValueChanged(b.ref.Value())
		}
	})
	return b.ref
}

var _ spot.Component = &ProgressIndicator{}
