//go:build !fltk && windows

package ui

import (
	"github.com/roblillack/spot"
	wui "github.com/rodrigocfd/windigo/ui"
	"github.com/rodrigocfd/windigo/win"
)

type nativeTypeButton = wui.Button

func (b *Button) callback() {
	if b.OnClick != nil {
		b.OnClick()
	}
}

func (b *Button) Update(nextControl spot.Control) bool {
	next, ok := nextControl.(*Button)
	if !ok {
		return false
	}

	if b.ref == nil {
		return false
	}

	if next.Title != b.Title {
		b.Title = next.Title
		b.ref.SetText(b.Title)
	}

	b.OnClick = next.OnClick
	return true
}

func (b *Button) Mount(parent spot.Control) any {
	if b.ref != nil {
		return b.ref
	}

	if parent == nil {
		return nil
	}

	window, ok := parent.(*Window)
	if !ok || window == nil || window.ref == nil {
		return nil
	}

	b.ref = wui.NewButton(window.ref, wui.ButtonOpts().
		Position(win.POINT{X: int32(b.X), Y: int32(b.Y)}).
		Size(win.SIZE{Cx: int32(b.Width), Cy: int32(b.Height)}).
		Text(b.Title))
	b.ref.On().BnClicked(b.callback)

	return b.ref
}

func (b *Button) Unmount() {
	// if b.ref == nil {
	// 	return
	// }

	// b.ref.Destroy()
	// b.ref = nil
	panic("not implemented")
}
