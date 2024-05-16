package main

import (
	"fmt"

	"github.com/roblillack/spot"
	"github.com/roblillack/spot/ui"
)

func main() {
	ui.Init()

	spot.MountFn(func(ctx *spot.RenderContext) spot.Component {
		counter, setCounter := spot.UseState[int](ctx, 0)

		buttonTitle := "Click me!"
		if counter > 0 {
			buttonTitle = fmt.Sprintf("Clicked %d times!", counter)
		}

		return &ui.Window{
			Title:  "Hello World!",
			Width:  200,
			Height: 125,
			Children: []spot.Component{
				&ui.Button{
					X: 25, Y: 50, Width: 150, Height: 25,
					Title: buttonTitle,
					OnClick: func() {
						setCounter(counter + 1)
					},
				},
			},
		}
	})

	ui.Run()
}
