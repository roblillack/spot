//go:build !cocoa && (fltk || !darwin)

package ui

import (
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
			// w.ref.SetSizeRange(w.Resizable)
		}
	}

	return true
}

func (w *Window) Mount(parent spot.Control) any {
	w.ref = goFltk.NewWindow(w.Width, w.Height, w.Title)
	// for _, child := range w.children {
	// 	w.mountChild(child)
	// }
	// w.ref.SetAllowsResizing(w.Resizable)

	// w.ref.MakeKeyAndOrderFront()
	// w.ref.AddDefaultQuitMenu()
	w.ref.End()
	w.ref.Show()
	return w.ref
}

func (w *Window) Unmount() {
	if w.ref != nil {
		w.ref.Hide()
		w.ref.Destroy()
		w.ref = nil
	}
}
