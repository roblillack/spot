//go:build !cocoa && (fltk || !darwin)

package ui

import (
	"image"
	"log"
	"slices"

	"github.com/llgcode/draw2d/draw2dimg"
	"github.com/llgcode/draw2d/draw2dkit"
	goFltk "github.com/pwiecz/go-fltk"
	"github.com/roblillack/spot"
	"github.com/roblillack/spot/ui/canvas"
)

type nativeTypeCanvas = *goFltk.Box

func (c *Canvas) Update(nextControl spot.Control) bool {
	next, ok := nextControl.(*Canvas)
	if !ok {
		return false
	}

	if c.ref == nil {
		return false
	}

	c.OnClick = next.OnClick
	// b.ref.SetCallback(b.OnClick)

	if !slices.Equal(c.Elements, next.Elements) {
		c.Elements = next.Elements
		c.draw()
	}

	return true
}

func (c *Canvas) Mount(parent spot.Control) any {
	if c.ref != nil {
		return c.ref
	}

	if parent == nil {
		return nil
	}

	c.ref = goFltk.NewBox(goFltk.DOWN_BOX, c.X, c.Y, c.Width, c.Height)
	// c.ref.SetCallback(c.callback)
	// c.ref.SetCallbackCondition(goFltk.WhenRelease)
	c.ref.SetEventHandler(c.handleEvent)
	c.draw()

	if window, ok := parent.(*Window); ok && window != nil && window.ref != nil {
		window.ref.Add(c.ref)
	}

	return c.ref
}

func (b *Canvas) Unmount() {
	if b.ref == nil {
		return
	}

	b.ref.Destroy()
	b.ref = nil
}

func (c *Canvas) handleEvent(e goFltk.Event) bool {
	if c.ref == nil || c.OnClick == nil {
		return false
	}

	if e == goFltk.PUSH || e == goFltk.DRAG {
		c.OnClick(goFltk.EventX(), goFltk.EventY())
		return true
	}

	return false
}

func (c *Canvas) draw() {
	if c.ref == nil {
		return
	}

	img := image.NewRGBA(image.Rect(0, 0, c.Width, c.Height))

	gc := draw2dimg.NewGraphicContext(img)

	for _, i := range c.Elements {
		switch e := i.(type) {
		case canvas.Circle:
			gc.MoveTo(float64(e.X), float64(e.Y))
			gc.SetFillColor(e.Fill)
			gc.SetStrokeColor(e.Stroke)
			gc.SetLineWidth(float64(e.StrokeWidth))
			draw2dkit.Circle(gc, float64(e.X), float64(e.Y), float64(e.Radius))
			gc.Fill()
			draw2dkit.Circle(gc, float64(e.X), float64(e.Y), float64(e.Radius))
			gc.Stroke()
		}
	}

	fimg, err := goFltk.NewRgbImageFromImage(img)
	if err != nil {
		log.Println(err)
		return
	}
	c.ref.SetImage(fimg)
	c.ref.Redraw()
}
