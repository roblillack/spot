//go:build !fltk && (darwin || cocoa)

package ui

import (
	"log"

	"github.com/roblillack/spot"
	"github.com/roblillack/spot/ui/internal/cocoa"
)

type nativeTypeWindow = *cocoa.Window

func (w *Window) Update(nextComponent spot.Control) bool {
	next, ok := nextComponent.(*Window)
	if !ok {
		return false
	}

	if next.Title != w.Title {
		w.Title = next.Title
		if w.ref != nil {
			w.ref.SetTitle(w.Title)
		}
	}

	if next.Resizable != w.Resizable {
		w.Resizable = next.Resizable
		if w.ref != nil {
			w.ref.SetAllowsResizing(w.Resizable)
		}
	}

	if next.Width != w.Width || next.Height != w.Height {
		w.Width = next.Width
		w.Height = next.Height
	}

	return true
}

func (w *Window) Mount(ctx *spot.RenderContext, parent spot.Control) any {
	w.ref = cocoa.NewCenteredWindow(w.Title, w.Width, w.Height)
	w.ref.SetAllowsResizing(w.Resizable)

	w.ref.MakeKeyAndOrderFront()
	w.ref.AddDefaultQuitMenu()
	w.ref.OnDidResize(func(wnd *cocoa.Window) {
		log.Printf("Window resized to %dx%d", w.ContentWidth(), w.ContentHeight())
		w.Width = w.ContentWidth()
		w.Height = w.ContentHeight()
		ctx.Layout()
	})
	return w.ref
}

func (w *Window) onResize(wnd *cocoa.Window) {
	log.Printf("Window resized to %dx%d", w.ContentWidth(), w.ContentHeight())
}

func (w *Window) ContentWidth() int {
	width, _ := w.ref.Size()
	return width
}

func (w *Window) ContentHeight() int {
	_, height := w.ref.Size()
	return height
}

func (c *Window) Layout(ctx *spot.RenderContext, parent spot.Control) {
	// No-op
}
