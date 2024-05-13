//go:build !cocoa && (fltk || !darwin)

package ui

import (
	goFltk "github.com/pwiecz/go-fltk"
	"github.com/roblillack/spot"
)

type Button struct {
	X       int
	Y       int
	Width   int
	Height  int
	Title   string
	OnClick func()
	ref     *goFltk.Button
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
		b.ref.SetLabel(b.Title)
	}

	b.OnClick = next.OnClick
	b.ref.SetCallback(b.OnClick)
	return true
}

func (b *Button) Mount() any {
	if b.ref != nil {
		return b.ref
	}

	b.ref = goFltk.NewButton(b.X, b.Y, b.Width, b.Height)
	b.ref.SetLabel(b.Title)
	b.ref.SetCallback(b.OnClick)
	return b.ref
}
