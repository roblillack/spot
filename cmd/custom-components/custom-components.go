package main

import (
	"fmt"
	"journey/cocoa"
	"journey/spot"

	"github.com/mojbro/gocoa"
)

func NoProps(ctx *spot.RenderContext) spot.Component {
	return &cocoa.Button{
		X: 25, Y: 30, Width: 150, Height: 25,
		Title:   "Custom button",
		OnClick: func() { fmt.Println("Button clicked!") },
	}
}

func CustomButton(ctx *spot.RenderContext, title string) spot.Component {
	return &cocoa.Button{
		X:      25,
		Y:      80,
		Width:  150,
		Height: 25,
		Title:  title,
		OnClick: func() {
			fmt.Println("Button clicked!")
		},
	}
}

func ButtonOrNot(ctx *spot.RenderContext, counter int) spot.Component {
	if counter < 3 {
		return nil
	}

	return &cocoa.Button{
		X:      25,
		Y:      30,
		Width:  150,
		Height: 25,
		Title:  "okokok",
		OnClick: func() {
			fmt.Println("Button clicked!")
		},
	}
}

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
				ctx.Make(func(x *spot.RenderContext) spot.Component {
					return CustomButton(x, buttonTitle)
				}),
				ctx.Make(NoProps),
				&cocoa.Button{
					X:      25,
					Y:      50,
					Width:  150,
					Height: 25,
					Title:  buttonTitle,
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
