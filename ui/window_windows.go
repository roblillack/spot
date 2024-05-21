//go:build !fltk && windows

package ui

import (
	"github.com/roblillack/spot"
	wui "github.com/rodrigocfd/windigo/ui"
	"github.com/rodrigocfd/windigo/win"
)

type nativeTypeWindow = wui.WindowMain

func (w *Window) Update(nextComponent spot.Control) bool {
	next, ok := nextComponent.(*Window)
	if !ok {
		return false
	}

	if next.Title != w.Title {
		w.Title = next.Title
		// if w.ref != nil {
		// 	w.ref.SetLabel(w.Title)
		// }
	}

	return true
}

func (w *Window) Mount(parent spot.Control) any {
	w.ref = wui.NewWindowMain(
		wui.WindowMainOpts().
			Title(w.Title).
			ClientArea(win.SIZE{Cx: int32(w.Width), Cy: int32(w.Height)}))

	activeWindow = w
	return w.ref
}
