//go:build !fltk && (darwin || cocoa)

package ui

import (
	"github.com/roblillack/gocoa"
	"github.com/roblillack/spot"
)

type nativeTypeCheckbox = *gocoa.Button

func (w *Checkbox) onClick() {
	w.Checked = !w.Checked
	if w.Checked {
		w.ref.SetState(gocoa.ButtonStateValueOn)
	} else {
		w.ref.SetState(gocoa.ButtonStateValueOff)
	}
	w.OnChange(w.Checked)
}

func (w *Checkbox) Update(nextComponent spot.Control) bool {
	next, ok := nextComponent.(*Checkbox)
	if !ok {
		return false
	}

	if w.ref == nil {
		return false
	}

	if next.Checked != w.Checked {
		w.Checked = next.Checked
		if w.Checked {
			w.ref.SetState(gocoa.ButtonStateValueOn)
		} else {
			w.ref.SetState(gocoa.ButtonStateValueOff)
		}
	}

	if next.Label != w.Label {
		w.Label = next.Label
		w.ref.SetTitle(w.Label)
	}

	w.OnChange = next.OnChange
	if w.OnChange == nil {
		w.ref.OnClick(nil)
	} else {
		w.ref.OnClick(w.onClick)
	}

	return true
}

func (w *Checkbox) Mount(parent spot.Control) any {
	if w.ref != nil {
		return w.ref
	}

	w.ref = gocoa.NewButton(w.X, w.Y, w.Width, w.Height)
	w.ref.SetTitle(w.Label)
	w.ref.SetButtonType(gocoa.ButtonTypeSwitch)
	if w.Checked {
		w.ref.SetState(gocoa.ButtonStateValueOn)
	} else {
		w.ref.SetState(gocoa.ButtonStateValueOff)
	}
	if w.OnChange == nil {
		w.ref.OnClick(nil)
	} else {
		w.ref.OnClick(w.onClick)
	}

	if window, ok := parent.(*Window); ok && window != nil && window.ref != nil {
		window.ref.AddButton(w.ref)
	}

	return w.ref
}
