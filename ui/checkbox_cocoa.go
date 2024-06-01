//go:build !fltk && (darwin || cocoa)

package ui

import (
	"github.com/roblillack/spot"
	"github.com/roblillack/spot/ui/internal/cocoa"
)

type nativeTypeCheckbox = *cocoa.Button

func (w *Checkbox) onClick() {
	w.Checked = !w.Checked
	if w.Checked {
		w.ref.SetState(cocoa.ButtonStateValueOn)
	} else {
		w.ref.SetState(cocoa.ButtonStateValueOff)
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
			w.ref.SetState(cocoa.ButtonStateValueOn)
		} else {
			w.ref.SetState(cocoa.ButtonStateValueOff)
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

func (w *Checkbox) Mount(ctx *spot.RenderContext, parent spot.Control) any {
	if w.ref != nil {
		return w.ref
	}

	w.ref = cocoa.NewButton(w.X, w.Y, w.Width, w.Height)
	w.ref.SetTitle(w.Label)
	w.ref.SetButtonType(cocoa.ButtonTypeSwitch)
	if w.Checked {
		w.ref.SetState(cocoa.ButtonStateValueOn)
	} else {
		w.ref.SetState(cocoa.ButtonStateValueOff)
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

func (c *Checkbox) Unmount() {
	if c.ref == nil {
		return
	}

	c.ref.Remove()
	c.ref = nil
}

func (c *Checkbox) Layout(ctx *spot.RenderContext, parent spot.Control) {
	x, y, w, h := calcLayout(parent, c.X, c.Y, c.Width, c.Height)
	c.ref.SetFrame(x, y, w, h)
}
