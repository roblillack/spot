//go:build !cocoa && (fltk || !darwin)

package ui

import (
	goFltk "github.com/pwiecz/go-fltk"
	"github.com/roblillack/spot"
)

type nativeTypeCheckbox = *goFltk.CheckButton

func (w *Checkbox) callback() {
	if w.OnChange != nil {
		w.OnChange(w.ref.Value())
	}
}

func (b *Checkbox) Update(nextComponent spot.Mountable) bool {
	next, ok := nextComponent.(*Checkbox)
	if !ok {
		return false
	}

	if b.ref == nil {
		return false
	}

	if next.Label != b.Label {
		b.Label = next.Label
		b.ref.SetLabel(b.Label)
	}

	if next.Checked != b.Checked {
		b.Checked = next.Checked
		b.ref.SetValue(b.Checked)
	}

	b.OnChange = next.OnChange
	if b.OnChange == nil {
		b.ref.SetCallback(nil)
	} else {
		b.ref.SetCallback(b.callback)
	}
	return true
}

func (b *Checkbox) Mount(ctx *spot.RenderContext, parent spot.Mountable) any {
	if b.ref != nil {
		return b.ref
	}

	x, y, w, h := CalcLayout(parent, b.X, b.Y, b.Width, b.Height)
	b.ref = goFltk.NewCheckButton(x, y, w, h)
	b.ref.SetLabel(b.Label)
	if b.OnChange == nil {
		b.ref.SetCallback(nil)
	} else {
		b.ref.SetCallback(b.callback)
	}
	b.ref.SetValue(b.Checked)

	if window, ok := parent.(*Window); ok && window != nil && window.ref != nil {
		window.ref.Add(b.ref)
	}

	return b.ref
}

func (b *Checkbox) Unmount() {
	if b.ref == nil {
		return
	}

	b.ref.Destroy()
	b.ref = nil
}

func (b *Checkbox) Layout(ctx *spot.RenderContext, parent spot.Container) {
	if b.ref == nil {
		return
	}

	x, y, w, h := CalcLayout(parent, b.X, b.Y, b.Width, b.Height)
	b.ref.Resize(x, y, w, h)
}
