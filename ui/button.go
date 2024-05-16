package ui

import "github.com/roblillack/spot"

var _ spot.HostComponent = &Button{}

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
