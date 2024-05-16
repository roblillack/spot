package main

import (
	"fmt"
	"time"

	"github.com/roblillack/spot"
	"github.com/roblillack/spot/ui"
)

func main() {
	ui.Init()

	root := spot.Render(spot.Make(func(ctx *spot.RenderContext) spot.Element {
		counter, setCounter := spot.UseState[int](ctx, 0)
		spot.UseEffect(ctx, func() {
			go func() {
				time.Sleep(3 * time.Second)
				setCounter(99)
			}()
		}, []any{})

		buttonTitle := "Click me!"
		if counter > 0 {
			buttonTitle = fmt.Sprintf("Clicked %d times!", counter)
		}

		return &ui.Window{
			Title:  "Hello World!",
			Width:  200,
			Height: 125,
			Children: []spot.Element{
				&ui.Button{
					X: 25, Y: 50, Width: 150, Height: 25,
					Title: buttonTitle,
					OnClick: func() {
						setCounter(counter + 1)
					},
				},
			},
		}
	}))
	root.Mount(nil)

	ui.Run()
}
