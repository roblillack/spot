//go:build !cocoa && (fltk || (!darwin && !windows))

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

func (b *Checkbox) Update(nextComponent spot.Control) bool {
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

func (b *Checkbox) Mount(parent spot.Control) any {
	if b.ref != nil {
		return b.ref
	}

	b.ref = goFltk.NewCheckButton(b.X, b.Y, b.Width, b.Height)
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
