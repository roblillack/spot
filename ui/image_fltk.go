//go:build !cocoa && (fltk || !darwin)

package ui

import (
	"log"

	goFltk "github.com/pwiecz/go-fltk"
	"github.com/roblillack/spot"
)

// There's currently no way to get the current RgbImage out of the
// box, so we have to keep a reference to it.
type fltkImageWrapper struct {
	box *goFltk.Box
	img *goFltk.RgbImage
}

type nativeTypeImage = *fltkImageWrapper

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

	box := goFltk.NewBox(goFltk.DOWN_BOX, c.X, c.Y, c.Width, c.Height)
	box.SetEventHandler(c.handleEvent)
	c.ref = &fltkImageWrapper{
		box: box,
	}
	c.draw()

	if window, ok := parent.(*Window); ok && window != nil && window.ref != nil {
		window.ref.Add(box)
	}

	return c.ref
}

func (b *Image) Unmount() {
	if b.ref == nil {
		return
	}

	if b.ref.img != nil {
		b.ref.img.Destroy()
	}

	if b.ref.box != nil {
		b.ref.box.Destroy()
	}

	b.ref = nil
}

func (c *Image) handleEvent(e goFltk.Event) bool {
	if c.ref == nil || c.OnClick == nil {
		return false
	}

	// if e == goFltk.PUSH || e == goFltk.DRAG {
	// 	c.OnClick(goFltk.EventX(), goFltk.EventY(), goFltk.EventButton1())
	// 	return true
	// }

	if e == goFltk.PUSH {
		c.OnClick(goFltk.EventX()-c.X, goFltk.EventY()-c.Y, !goFltk.EventButton1())
	}

	return false
}

func (c *Image) draw() {
	if c.ref == nil || isImageNil(c.Image) {
		return
	}

	fimg, err := goFltk.NewRgbImageFromImage(c.Image)
	if err != nil {
		log.Println(err)
		return
	}
	oldImg := c.ref.img
	c.ref.img = fimg
	c.ref.box.SetImage(fimg)
	c.ref.box.Redraw()
	if oldImg != nil {
		oldImg.Destroy()
	}
}
