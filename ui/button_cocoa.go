//go:build !fltk && (darwin || cocoa)

package ui

import (
	"github.com/roblillack/gocoa"
	"github.com/roblillack/spot"
)

type nativeTypeButton = *gocoa.Button

func (b *Button) Update(nextComponent spot.Control) bool {
	next, ok := nextComponent.(*Button)
	if !ok {
		return false
	}

	if b.ref == nil {
		return false
	}

	if next.Title != b.Title {
		b.Title = next.Title
		b.ref.SetTitle(b.Title)
	}

	b.OnClick = next.OnClick
	b.ref.OnClick(b.OnClick)
	return true
}

func (b *Button) Mount(parent spot.Control) any {
	if b.ref != nil {
		return b.ref
	}

	b.ref = gocoa.NewButton(b.X, b.Y, b.Width, b.Height)
	b.ref.SetTitle(b.Title)
	b.ref.OnClick(b.OnClick)

	if window, ok := parent.(*Window); ok && window != nil && window.ref != nil {
		window.ref.AddButton(b.ref)
	}

	return b.ref
}

func (b *Button) Unmount() {
	panic("not implemented")
}
