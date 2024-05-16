//go:build !cocoa && (fltk || !darwin)

package ui

import (
	goFltk "github.com/pwiecz/go-fltk"
	"github.com/roblillack/spot"
)

type Slider struct {
	X              int
	Y              int
	Width          int
	Height         int
	Min            float64
	Max            float64
	Value          float64
	Type           goFltk.SliderType
	OnValueChanged func(float64)
	ref            *goFltk.Slider
}

func (b *Slider) Equals(other spot.Control) bool {
	next, ok := other.(*Slider)
	if !ok {
		return false
	}

	if b == nil && next != nil || b != nil && next == nil {
		return false
	}

	return next.Max == b.Max && next.Min == b.Min &&
		next.Value == b.Value &&
		next.Type == b.Type
}

func (b *Slider) Update(nextComponent spot.Control) bool {
	next, ok := nextComponent.(*Slider)
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

	if next.Type != b.Type {
		b.Type = next.Type
		b.ref.SetType(b.Type)
	}

	return true
}

func (b *Slider) Mount(parent spot.Control) any {
	if b.ref != nil {
		return b.ref
	}

	b.ref = goFltk.NewSlider(b.X, b.Y, b.Width, b.Height)
	b.ref.SetMaximum(b.Max)
	b.ref.SetMinimum(b.Min)
	b.ref.SetValue(b.Value)
	// b.ref.SetType(b.Type)
	// b.ref.SetType(goFltk.HOR_SLIDER)
	b.ref.SetType(goFltk.HOR_NICE_SLIDER)
	b.ref.SetBox(goFltk.FLAT_BOX)
	b.ref.SetCallback(func() {
		if b.OnValueChanged != nil {
			b.OnValueChanged(b.ref.Value())
		}
	})

	if window, ok := parent.(*Window); ok && window != nil && window.ref != nil {
		window.ref.Add(b.ref)
	}

	return b.ref
}

var _ spot.Control = &Slider{}
