//go:build !fltk && (darwin || cocoa)

package ui

import (
	"github.com/roblillack/spot"
)

type nativeTypeImage = struct{}

func (b *Image) Update(nextComponent spot.Control) bool {
	panic("not implemented")
}

func (b *Image) Mount(parent spot.Control) any {
	panic("not implemented")
}

func (b *Image) Unmount() {
	panic("not implemented")
}
