package main

import (
	"fmt"
	"journey/cocoa"
	"journey/spot"

	"github.com/mojbro/gocoa"
)

func main() {
	// runtime.LockOSThread()
	gocoa.InitApplication()

	root := spot.Make(func(ctx *spot.RenderContext) spot.Component {
		counter, setCounter := spot.UseState[int](ctx, 0)

		buttonTitle := "Click me!"
		if counter > 0 {
			buttonTitle = fmt.Sprintf("Clicked %d times!", counter)
		}

		return &cocoa.Window{
			Title:  "Hello World!",
			Width:  200,
			Height: 125,
			Children: []spot.Component{
				&cocoa.Button{
					X: 25, Y: 50, Width: 150, Height: 25,
					Title: buttonTitle,
					OnClick: func() {
						setCounter(counter + 1)
					},
				},
			},
		}
	})

	root.Mount()

	gocoa.RunApplication()
}
