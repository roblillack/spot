//go:build !fltk && (darwin || cocoa)

package ui

import (
	"journey/spot"

	"github.com/mojbro/gocoa"
)

type Button struct {
	X       int
	Y       int
	Width   int
	Height  int
	Title   string
	OnClick func()
	ref     *gocoa.Button
}

func (b *Button) Update(nextComponent spot.Component) bool {
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

func (b *Button) Mount() any {
	if b.ref != nil {
		return b.ref
	}

	b.ref = gocoa.NewButton(b.X, b.Y, b.Width, b.Height)
	b.ref.SetTitle(b.Title)
	b.ref.OnClick(b.OnClick)
	return b.ref
}

var _ spot.Component = &Button{}
