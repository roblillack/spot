package ui

import "github.com/roblillack/spot"

var _ spot.Component = &Window{}
var _ spot.Container = &Window{}

func (w *Window) GetChildren() []spot.Component {
	return w.Children
}
