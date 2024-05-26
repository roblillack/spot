package main

import (
	"image/color"
	"log"
	"math/rand"

	"github.com/roblillack/spot"
	"github.com/roblillack/spot/ui"
	"github.com/roblillack/spot/ui/canvas"
)

func RandomColor() color.Color {
	return color.RGBA{uint8(rand.Intn(256)), uint8(rand.Intn(256)), uint8(rand.Intn(256)), 0xff}
}

func main() {
	ui.Init()

	spot.MountFn(func(ctx *spot.RenderContext) spot.Component {
		circles, setCircles := spot.UseState(ctx, []canvas.Element{})

		log.Printf("Rendering %d elements\n", len(circles))

		return &ui.Window{
			Title:  "Spot Draw",
			Width:  800,
			Height: 600,
			Children: []spot.Component{
				&ui.Canvas{
					X: 10, Y: 10, Width: 780, Height: 580,
					Elements: circles,
					OnClick: func(x, y int) {
						log.Printf("Clicked %d, %d\n", x, y)
						setCircles(append(circles, canvas.Circle{
							X: x, Y: y,
							Radius:      10,
							Fill:        RandomColor(),
							Stroke:      RandomColor(),
							StrokeWidth: 2,
						}))
					},
				},
			},
		}
	})

	ui.Run()
}
