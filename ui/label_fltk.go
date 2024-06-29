//go:build !cocoa && (fltk || !darwin)

package ui

import (
	goFltk "github.com/pwiecz/go-fltk"
	"github.com/roblillack/spot"
)

type nativeTypeLabel = *goFltk.Box

func (c *Label) Update(nextComponent spot.Mountable) bool {
	next, ok := nextComponent.(*Label)
	if !ok {
		return false
	}

	if c.ref == nil {
		return false
	}

	if next.Value != c.Value {
		c.Value = next.Value
		if c.ref != nil {
			// if w.ref.Buffer() == nil {
			// 	w.ref.SetBuffer(goFltk.NewTextBuffer())
			// }
			// w.ref.Buffer().SetText(w.Value)
			c.ref.SetLabel(c.Value)
		}
	}

	if next.FontSize != c.FontSize {
		c.FontSize = next.FontSize
		if c.ref != nil && c.FontSize > 0 {
			c.ref.SetLabelSize(c.FontSize)
		}
	}

	if next.Align != c.Align {
		c.setAlign(next.Align)
	}

	return true
}

func (c *Label) setAlign(a LabelAlignment) {
	c.Align = a
	if c.ref == nil {
		return
	}
	switch a {
	case LabelAlignmentLeft:
		c.ref.SetAlign(goFltk.ALIGN_INSIDE | goFltk.ALIGN_LEFT)
	case LabelAlignmentCenter:
		c.ref.SetAlign(goFltk.ALIGN_INSIDE | goFltk.ALIGN_CENTER)
	case LabelAlignmentRight:
		c.ref.SetAlign(goFltk.ALIGN_INSIDE | goFltk.ALIGN_RIGHT)
	}
}

func (c *Label) Mount(ctx *spot.RenderContext, parent spot.Mountable) any {
	if c.ref != nil {
		return c.ref
	}

	// w.ref = goFltk.NewTextDisplay(w.X, w.Y, w.Width, w.Height)
	x, y, w, h := CalcLayout(parent, c.X, c.Y, c.Width, c.Height)
	c.ref = goFltk.NewBox(goFltk.NO_BOX, x, y, w, h)
	c.ref.SetLabel(c.Value)
	c.setAlign(c.Align)
	// buf := goFltk.NewTextBuffer()
	// buf.SetText(w.Value)
	// w.ref.SetBuffer(buf)
	if c.FontSize > 0 {
		// w.ref.SetTextSize(w.FontSize)
		c.ref.SetLabelSize(c.FontSize)
	}
	// w.ref.HideCursor()

	if window, ok := parent.(*Window); ok && window != nil && window.ref != nil {
		window.ref.Add(c.ref)
	}

	return c.ref
}

func (c *Label) Unmount() {
	if c.ref != nil {
		c.ref.Destroy()
		c.ref = nil
	}
}

func (c *Label) Layout(ctx *spot.RenderContext, parent spot.Container) {
	x, y, w, h := CalcLayout(parent, c.X, c.Y, c.Width, c.Height)
	c.ref.Resize(x, y, w, h)
}
