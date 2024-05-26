package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"log"

	"github.com/roblillack/spot"
	"github.com/roblillack/spot/ui"
)

type Circle struct {
	X, Y int
}

const Radius = 20

type State struct {
	Circles []Circle
	Active  *Circle
}

type History[T any] struct {
	Data    []T
	Current int
}

func (s State) OnClick(x, y int, secondary bool) State {
	log.Printf("Clicked %d, %d\n", x, y)

	for _, i := range s.Circles {
		for cy := -Radius; cy <= Radius; cy++ {
			for cx := -Radius; cx <= Radius; cx++ {
				if cx*cx+cy*cy <= Radius*Radius {
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
			X: x, Y: y,
		}),
		Active: nil,
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
		drawCircle(img, i.X, i.Y, 20, color.Black)
		if state.Active != nil && state.Active.X == i.X && state.Active.Y == i.Y {
			drawCircle(img, i.X, i.Y, 18, color.RGBA{0xff, 0x00, 0x00, 0xff})
		} else {
			drawCircle(img, i.X, i.Y, 18, color.White)
		}
	}

	return img
}

func main() {
	ui.Init()

	spot.MountFn(func(ctx *spot.RenderContext) spot.Component {
		history, setHistory := spot.UseState(ctx, History[State]{Data: []State{State{}}, Current: 0})
		state := history.Data[history.Current]

		img := renderImg(780, 550, state)

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
				&ui.Image{
					X: 10, Y: 40, Width: 780, Height: 550,
					Image: img,
					OnClick: func(x, y int, secondary bool) {
						newState := state.OnClick(x, y, secondary)
						setHistory(History[State]{
							Data:    append(history.Data[0:history.Current+1], newState),
							Current: history.Current + 1,
						})
					},
				},
			},
		}
	})

	ui.Run()
}
