//go:build !fltk && (darwin || cocoa)

package ui

import (
	"fmt"
	"image"

	"github.com/roblillack/gocoa"
	"github.com/roblillack/spot"
)

type nativeTypeImage = *gocoa.CustomButton

func (c *Image) Update(nextControl spot.Control) bool {
	next, ok := nextControl.(*Image)
	if !ok {
		return false
	}

	if c.ref == nil {
		return false
	}

	c.OnClick = next.OnClick
	c.Image = next.Image
	c.draw()

	return true
}

func (c *Image) Mount(parent spot.Control) any {
	if c.ref != nil {
		return c.ref
	}

	if parent == nil {
		return nil
	}

	c.ref = gocoa.NewCustomButton(c.X, c.Y, c.Width, c.Height)
	c.ref.OnClick(c.handleClick)
	c.draw()

	if window, ok := parent.(*Window); ok && window != nil && window.ref != nil {
		window.ref.AddCustomButton(c.ref)
	}

	return c.ref
}

func (c *Image) Unmount() {
	if c.ref == nil {
		return
	}

	c.ref.Remove()
	c.ref = nil
}

func (c *Image) draw() {
	if c.Image == nil {
		return
	}

	switch img := c.Image.(type) {
	case *image.RGBA:
		c.ref.SetImage(img)
	default:
		panic(fmt.Sprintf("unsupported image type: %T", img))
	}
}

func (c *Image) handleClick(x, y int, secondary bool) {
	if c.OnClick != nil {
		c.OnClick(x, y, secondary)
	}
}
