package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"log"
	"slices"

	"github.com/roblillack/spot"
	"github.com/roblillack/spot/ui"
)

type Circle struct {
	X, Y int
	Size int
}

const DefaultRadius = 20

type History[T any] struct {
	Data    []T
	Current int
}

func NewHistory[T any](value T) History[T] {
	return History[T]{
		Data:    []T{value},
		Current: 0,
	}
}

func (h History[T]) CanUndo() bool {
	return h.Current > 0
}

func (h History[T]) CanRedo() bool {
	return h.Current < len(h.Data)-1
}

func (h History[T]) Get() T {
	return h.Data[h.Current]
}

func (h History[T]) Put(value T) History[T] {
	return History[T]{
		Data:    append(h.Data[0:h.Current+1], value),
		Current: h.Current + 1,
	}
}

type State struct {
	Circles []Circle
	Active  *Circle
}

func (s State) OnClick(x, y int, secondary bool) State {
	log.Printf("Clicked %d, %d\n", x, y)

	for _, i := range s.Circles {
		for cy := -i.Size; cy <= i.Size; cy++ {
			for cx := -i.Size; cx <= i.Size; cx++ {
				if cx*cx+cy*cy <= i.Size*i.Size {
					if i.X+cx == x && i.Y+cy == y {
						return State{
							Circles: s.Circles,
							Active:  &i,
						}
					}
				}
			}
		}
	}

	return State{
		Circles: append(s.Circles, Circle{
			X: x, Y: y, Size: DefaultRadius,
		}),
		Active: nil,
	}
}

func (s State) Resize(newSize int) State {
	if s.Active == nil {
		return s
	}

	idx := slices.Index(s.Circles, *s.Active)
	if idx == -1 {
		return s
	}

	newCircles := slices.Clone(s.Circles)
	resized := Circle{
		X:    s.Active.X,
		Y:    s.Active.Y,
		Size: newSize,
	}
	newCircles[idx] = resized

	return State{
		Circles: newCircles,
		Active:  &resized,
	}
}

func drawCircle(img *image.RGBA, cx, cy, r int, col color.Color) {
	for y := -r; y <= r; y++ {
		for x := -r; x <= r; x++ {
			if x*x+y*y <= r*r {
				img.Set(cx+x, cy+y, col)
			}
		}
	}
}

func renderImg(w, h int, state State) image.Image {
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.White},
		image.Point{}, draw.Src)

	for _, i := range state.Circles {
		drawCircle(img, i.X, i.Y, i.Size, color.Black)
		if state.Active != nil && state.Active.X == i.X && state.Active.Y == i.Y {
			drawCircle(img, i.X, i.Y, i.Size-2, color.RGBA{0xff, 0x00, 0x00, 0xff})
		} else {
			drawCircle(img, i.X, i.Y, i.Size-2, color.White)
		}
	}

	return img
}

func main() {
	ui.Init()

	spot.MountFn(func(ctx *spot.RenderContext) spot.Component {
		history, setHistory := spot.UseState(ctx, NewHistory(State{}))
		state := history.Get()

		img := renderImg(780, 550, state)

		var sizeSlider spot.Fragment
		if state.Active != nil {
			sizeSlider = spot.Fragment{
				&ui.Label{
					X: 400, Y: 10, Width: 80, Height: 25,
					Value: "Size:",
				},
				&ui.Slider{
					X: 480, Y: 10, Width: 200, Height: 25,
					Min: 10, Max: 100,
					Value: float64(state.Active.Size),
					OnValueChanged: func(value float64) {
						setHistory(history.Put(state.Resize(int(value))))
					},
				},
			}
		}

		return &ui.Window{
			Title:  "Spot Draw",
			Width:  800,
			Height: 600,
			Children: []spot.Component{
				&ui.Button{
					X: 10, Y: 10, Width: 80, Height: 25,
					Title: "Undo",
					OnClick: func() {
						if history.Current > 0 {
							setHistory(History[State]{Data: history.Data, Current: history.Current - 1})
						}
					},
				},
				&ui.Button{
					X: 100, Y: 10, Width: 80, Height: 25,
					Title: "Redo",
					OnClick: func() {
						if history.Current < len(history.Data)-1 {
							setHistory(History[State]{Data: history.Data, Current: history.Current + 1})
						}
					},
				},
				&ui.Label{
					X: 190, Y: 10, Width: 200, Height: 25,
					Value: fmt.Sprintf("History: %d/%d", history.Current+1, len(history.Data)),
				},
				sizeSlider,
				&ui.Image{
					X: 10, Y: 40, Width: 780, Height: 550,
					Image: img,
					OnClick: func(x, y int, secondary bool) {
						setHistory(history.Put(state.OnClick(x, y, secondary)))
					},
				},
			},
		}
	})

	ui.Run()
}
