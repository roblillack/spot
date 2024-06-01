//go:build !fltk && (darwin || cocoa)

package ui

import (
	"github.com/roblillack/spot"
	"github.com/roblillack/spot/ui/internal/cocoa"
)

type nativeTypeTextField = *cocoa.TextField

func (w *TextField) Update(nextComponent spot.Control) bool {
	next, ok := nextComponent.(*TextField)
	if !ok {
		return false
	}

	if next.Value != w.Value {
		w.Value = next.Value
		if w.ref != nil {
			w.ref.SetStringValue(w.Value)
		}
	}

	// if next.Editable != w.Editable {
	// 	w.Editable = next.Editable
	// 	if w.ref != nil {
	// 		w.ref.SetEditable(w.Editable)
	// 	}
	// }

	// if next.Bezeled != w.Bezeled {
	// 	w.Bezeled = next.Bezeled
	// 	if w.ref != nil {
	// 		w.ref.SetBezeled(w.Bezeled)
	// 	}
	// }

	// if next.Selectable != w.Selectable {
	// 	w.Selectable = next.Selectable
	// 	if w.ref != nil {
	// 		w.ref.SetSelectable(w.Selectable)
	// 	}
	// }

	// if next.FontSize != w.FontSize {
	// 	w.FontSize = next.FontSize
	// 	if w.ref != nil {
	// 		w.ref.SetFontSize(w.FontSize)
	// 	}
	// }

	// if next.NoBackground != w.NoBackground {
	// 	w.NoBackground = next.NoBackground
	// 	if w.ref != nil {
	// 		w.ref.SetDrawsBackground(!w.NoBackground)
	// 	}
	// }

	w.OnChange = next.OnChange
	w.ref.OnChange(w.OnChange)

	return true
}

func (c *TextField) Mount(ctx *spot.RenderContext, parent spot.Control) any {
	if c.ref != nil {
		return c.ref
	}

	x, y, w, h := calcLayout(parent, c.X, c.Y, c.Width, c.Height)
	c.ref = cocoa.NewTextField(x, y, w, h)
	// w.ref.SetEditable(w.Editable)
	// w.ref.SetBezeled(w.Bezeled)
	// w.ref.SetSelectable(w.Selectable)
	c.ref.SetStringValue(c.Value)
	c.ref.SetFontFamily("Arial")
	c.ref.OnChange(c.OnChange)
	// w.ref.SetFontSize(w.FontSize)
	// w.ref.SetDrawsBackground(!w.NoBackground)

	if window, ok := parent.(*Window); ok && window != nil && window.ref != nil {
		window.ref.AddTextField(c.ref)
	}

	return c.ref
}

func (c *TextField) Unmount() {
	if c.ref == nil {
		return
	}

	c.ref.Remove()
	c.ref = nil
}

func (c *TextField) Layout(ctx *spot.RenderContext, parent spot.Control) {
	x, y, w, h := calcLayout(parent, c.X, c.Y, c.Width, c.Height)
	c.ref.SetFrame(x, y, w, h)
}
