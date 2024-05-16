package ui

import "github.com/roblillack/spot"

type Checkbox struct {
	X        int
	Y        int
	Width    int
	Height   int
	Label    string
	Checked  bool
	OnChange func(checked bool)
	ref      nativeTypeCheckbox
}

var _ spot.Component = &Checkbox{}
var _ spot.Control = &Checkbox{}

func (c *Checkbox) Render(ctx *spot.RenderContext) spot.Component {
	return c
}
