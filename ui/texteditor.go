package ui

import "github.com/roblillack/spot"

type TextEditor struct {
	X        int
	Y        int
	Width    int
	Height   int
	Text     string
	FontSize int
	OnChange func(content string)
	ref      nativeTypeTextEditor
}

var _ spot.Component = &TextEditor{}
var _ spot.Control = &TextEditor{}

func (c *TextEditor) Render(ctx *spot.RenderContext) spot.Component {
	return c
}
