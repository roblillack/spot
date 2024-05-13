package fltk

import (
	"journey/spot"

	goFltk "github.com/pwiecz/go-fltk"
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
	ref            *goFltk.Spinner
}

func (b *Spinner) Equals(other spot.Component) bool {
	next, ok := other.(*Spinner)
	if !ok {
		return false
	}

	if b == nil && next != nil || b != nil && next == nil {
		return false
	}

	return next.Max == b.Max && next.Min == b.Min &&
		next.Value == b.Value &&
		next.Step == b.Step
}

func (b *Spinner) Update(nextComponent spot.Component) bool {
	next, ok := nextComponent.(*Spinner)
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

	if next.Step != b.Step {
		b.Step = next.Step
		b.ref.SetStep(b.Step)
	}

	return true
}

func (b *Spinner) Mount() any {
	if b.ref != nil {
		return b.ref
	}

	b.ref = goFltk.NewSpinner(b.X, b.Y, b.Width, b.Height)
	b.ref.SetMaximum(b.Max)
	b.ref.SetMinimum(b.Min)
	b.ref.SetValue(b.Value)
	b.ref.SetStep(b.Step)
	b.ref.SetCallback(func() {
		if b.OnValueChanged != nil {
			b.OnValueChanged(b.ref.Value())
		}
	})
	return b.ref
}

var _ spot.Component = &Spinner{}
