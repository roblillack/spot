package main

import (
	"fmt"
	"journey/fltk"
	"journey/spot"
	"time"

	goFltk "github.com/pwiecz/go-fltk"
)

func main() {
	// runtime.LockOSThread()
	// gocoa.InitApplication()

	root := spot.Make(func(ctx *spot.RenderContext) spot.Component {
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

		return &fltk.Window{
			Title:  "Hello World!",
			Width:  200,
			Height: 125,
			Children: []spot.Component{
				&fltk.Button{
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

	// gocoa.RunApplication()
	goFltk.Run()
}
