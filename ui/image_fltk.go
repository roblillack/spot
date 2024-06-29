//go:build !cocoa && (fltk || !darwin)

package ui

import (
	"log"

	goFltk "github.com/pwiecz/go-fltk"
	"github.com/roblillack/spot"
)

type nativeTypeImage = *goFltk.Box

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

	c.ref = goFltk.NewBox(goFltk.DOWN_BOX, c.X, c.Y, c.Width, c.Height)
	c.ref.SetEventHandler(c.handleEvent)
	c.draw()

	if window, ok := parent.(*Window); ok && window != nil && window.ref != nil {
		window.ref.Add(c.ref)
	}

	return c.ref
}

func (b *Image) Unmount() {
	if b.ref == nil {
		return
	}

	b.ref.Destroy()
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
		c.OnClick(goFltk.EventX()-c.X, goFltk.EventY()-c.Y, goFltk.EventButton1())
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
	c.ref.SetImage(fimg)
	c.ref.Redraw()
}
