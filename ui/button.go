package ui

import "github.com/roblillack/spot"

type Button struct {
	X       int
	Y       int
	Width   int
	Height  int
	Title   string
	OnClick func()

	ref nativeTypeButton
}

var _ spot.HostComponent = &Button{}

func (b *Button) Render(ctx *spot.RenderContext) spot.Component {
	return b
}

func (b *Button) Equals(other spot.HostComponent) bool {
	next, ok := other.(*Button)
	if !ok {
		return false
	}

	if b == nil && next != nil || b != nil && next == nil {
		return false
	}

	return next.Title == b.Title
}
