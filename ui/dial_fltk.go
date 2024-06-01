//go:build !cocoa && (fltk || !darwin)

package ui

import (
	goFltk "github.com/pwiecz/go-fltk"
	"github.com/roblillack/spot"
)

type nativeTypeDial = *goFltk.Slider

func (b *Dial) Update(nextComponent spot.Control) bool {
	next, ok := nextComponent.(*Dial)
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

func (b *Dial) Mount(ctx *spot.RenderContext, parent spot.Control) any {
	if b.ref != nil {
		return b.ref
	}

	x, y, w, h := calcLayout(parent, b.X, b.Y, b.Width, b.Height)
	b.ref = goFltk.NewSlider(x, y, w, h)
	b.ref.SetMaximum(b.Max)
	b.ref.SetMinimum(b.Min)
	b.ref.SetValue(b.Value)
	// b.ref.SetType(b.Type)
	b.ref.SetType(goFltk.HOR_SLIDER)
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

func (b *Dial) Unmount() {
	if b.ref == nil {
		return
	}

	b.ref.Destroy()
	b.ref = nil
}

func (b *Dial) Layout(ctx *spot.RenderContext, parent spot.Control) {
	if b.ref == nil {
		return
	}

	x, y, w, h := calcLayout(parent, b.X, b.Y, b.Width, b.Height)
	b.ref.Resize(x, y, w, h)
}
