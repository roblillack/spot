//go:build !cocoa && (fltk || !darwin)

package ui

import (
	"log"

	goFltk "github.com/pwiecz/go-fltk"
	"github.com/roblillack/spot"
)

type nativeTypeWindow = *goFltk.Window

func (w *Window) Update(nextComponent spot.Control) bool {
	next, ok := nextComponent.(*Window)
	if !ok {
		return false
	}

	if next.Title != w.Title {
		w.Title = next.Title
		if w.ref != nil {
			w.ref.SetLabel(w.Title)
		}
	}

	if next.Width != w.Width || next.Height != w.Height {
		w.Width = next.Width
		w.Height = next.Height
	}

	if next.Resizable != w.Resizable {
		w.Resizable = next.Resizable
		if w.ref != nil {
			if w.Resizable {
				w.ref.SetSizeRange(1, 1, 0, 0, 0, 0, false)
			} else {
				w.ref.SetSizeRange(w.Width, w.Width, w.Height, w.Height, 0, 0, false)
			}
		}
	}

	return true
}

func (w *Window) Mount(ctx *spot.RenderContext, parent spot.Control) any {
	w.ref = goFltk.NewWindow(w.Width, w.Height, w.Title)
	// for _, child := range w.children {
	// 	w.mountChild(child)
	// }
	// w.ref.SetAllowsResizing(w.Resizable)

	// w.ref.MakeKeyAndOrderFront()
	// w.ref.AddDefaultQuitMenu()

	if w.Resizable {
		w.ref.SetSizeRange(1, 1, 0, 0, 0, 0, false)
	} else {
		w.ref.SetSizeRange(w.Width, w.Width, w.Height, w.Height, 0, 0, false)
	}
	w.ref.End()
	w.ref.Show()
	w.ref.SetResizeHandler(func() {
		log.Printf("Window resized to %dx%d", w.ContentWidth(), w.ContentHeight())
		w.Width = w.ContentWidth()
		w.Height = w.ContentHeight()
		ctx.Layout()
	})
	return w.ref
}

func (w *Window) ContentWidth() int {
	return w.ref.W()
}

func (w *Window) ContentHeight() int {
	return w.ref.H()
}

func (w *Window) Layout(ctx *spot.RenderContext, parent spot.Control) {
	// no-op
}
