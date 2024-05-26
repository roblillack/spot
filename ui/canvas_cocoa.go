//go:build !fltk && (darwin || cocoa)

package ui

import (
	"github.com/roblillack/spot"
)

type nativeTypeCanvas = struct{}

func (b *Canvas) Update(nextComponent spot.Control) bool {
	panic("not implemented")
}

func (b *Canvas) Mount(parent spot.Control) any {
	panic("not implemented")
}

func (b *Canvas) Unmount() {
	panic("not implemented")
}
