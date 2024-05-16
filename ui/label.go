package ui

import (
	"github.com/roblillack/spot"
)

type Label struct {
	X        int
	Y        int
	Width    int
	Height   int
	Value    string
	FontSize int
	ref      nativeTypeLabel
}

var _ spot.HostComponent = &Label{}

func (w *Label) Equals(other spot.HostComponent) bool {
	next, ok := other.(*Label)
	if !ok {
		return false
	}

	if w == nil && next != nil || w != nil && next == nil {
		return false
	}

	return next.Value == w.Value && w.FontSize == next.FontSize
}
