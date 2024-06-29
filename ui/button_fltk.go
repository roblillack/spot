//go:build !cocoa && (fltk || !darwin)

package ui

import (
	"log"

	goFltk "github.com/pwiecz/go-fltk"
	"github.com/roblillack/spot"
)

type nativeTypeButton = *goFltk.Button

func (b *Button) Update(nextControl spot.Mountable) bool {
	next, ok := nextControl.(*Button)
	if !ok {
		return false
	}

	if b.ref == nil {
		return false
	}

	if next.Title != b.Title {
		b.Title = next.Title
		b.ref.SetLabel(b.Title)
	}

	b.OnClick = next.OnClick
	b.ref.SetCallback(b.OnClick)
	return true
}

func (b *Button) Mount(ctx *spot.RenderContext, parent spot.Mountable) any {
	if b.ref != nil {
		return b.ref
	}

	if parent == nil {
		return nil
	}

	x, y, w, h := CalcLayout(parent, b.X, b.Y, b.Width, b.Height)
	b.ref = goFltk.NewButton(x, y, w, h)
	b.ref.SetLabel(b.Title)
	b.ref.SetCallback(b.OnClick)

	log.Printf("Button.Mount: parent=%T\n", parent)

	if container, ok := parent.(spot.Container); ok && container != nil {
		// container.MountChild(b)
	}

	if window, ok := parent.(*Window); ok && window != nil && window.ref != nil {
		window.ref.Add(b.ref)
	}

	return b.ref
}

func (b *Button) Unmount() {
	if b.ref == nil {
		return
	}

	b.ref.Destroy()
	b.ref = nil
}

func (b *Button) Layout(ctx *spot.RenderContext, parent spot.Container) {
	if b.ref == nil {
		return
	}

	x, y, w, h := CalcLayout(parent, b.X, b.Y, b.Width, b.Height)
	b.ref.Resize(x, y, w, h)
}
