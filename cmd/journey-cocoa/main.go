package main

import (
	"fmt"
	"journey/cocoa"
	"journey/spot"
	"math/rand"
	"strings"
	"time"

	"github.com/mojbro/gocoa"
)

// func Make2(render func(ctx *react.RenderContext) react.Component) react.Component {
// 	ctx := &react.RenderContext{
// 		render: render,
// 		values: make(map[int]any),
// 	}
// 	root := render(ctx)
// 	ctx.root = root
// 	// root.Build()
// 	// root.Update(root)
// 	return root
// }

func QuitButton(ctx *spot.RenderContext) spot.Component {
	enabled, setEnabled := spot.UseState(ctx, false)
	spot.UseEffect(ctx, func() {
		fmt.Println("Setting up timer!")
		go func() {
			time.Sleep(5 * time.Second)
			fmt.Println("Enabling button!")
			setEnabled(true)
		}()
	}, []any{})

	if !enabled {
		return &cocoa.Button{
			X:      210,
			Y:      100,
			Width:  180,
			Height: 25,
			Title:  "Please wait...",
		}
	}

	return &cocoa.Button{
		X:      200,
		Y:      100,
		Width:  180,
		Height: 25,
		Title:  "Quit",
		OnClick: func() {
			gocoa.TerminateApplication()
		},
	}
}

// https://stackoverflow.com/a/31832326
func RandStringBytesMaskImprSrcSB(n int) string {
	var src = rand.NewSource(time.Now().UnixNano())
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const (
		letterIdxBits = 6                    // 6 bits to represent a letter index
		letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
		letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
	)

	sb := strings.Builder{}
	sb.Grow(n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			sb.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return sb.String()
}

func main() {
	// runtime.LockOSThread()
	gocoa.InitApplication()
	// gocoa.OnApplicationDidFinishLaunching(func() {
	// 	fmt.Println("App running!")
	// })

	startTime := time.Now()
	spot.Make(func(ctx *spot.RenderContext) spot.Component {
		counter, setCounter := spot.UseState[int](ctx, 0)
		duration, setDuration := spot.UseState[time.Duration](ctx, 0.0)
		spot.UseEffect(ctx, func() {
			go func() {
				time.Sleep(50 * time.Millisecond)
				setDuration(time.Now().Sub(startTime))
			}()
		}, []any{duration})

		buttonTitle := "Click me!"
		if counter > 0 {
			buttonTitle = fmt.Sprintf("Clicked %d times!", counter)
		}

		// var but Component = nil
		// var fn func()
		// if counter > 5 {
		// 	fn = ctx.Wrap(func() {
		// 		gocoa.TerminateApplication()
		// 	})
		// }
		// if counter > 0 {
		// 	but = &Button{
		// 		X:       10,
		// 		Y:       10,
		// 		Width:   30,
		// 		Height:  30,
		// 		Title:   "X",
		// 		OnClick: fn,
		// 	}
		// }

		return &cocoa.Window{
			Title: "Hello World!",
			Width: 400, Height: 500,
			Children: []spot.Component{
				&cocoa.TextField{
					X: 10, Y: 410, Width: 380, Height: 80,
					Value: fmt.Sprintf("%02d:%02d.%03d", int(duration.Minutes())%60,
						int(duration.Seconds())%60, duration.Milliseconds()%1000),
					FontSize: 80,
					Editable: false, Selectable: false, Bezeled: false, NoBackground: false,
				},
				spot.Range(ctx, 0, 10, func(ctx *spot.RenderContext, i int) spot.Component {
					max := float64(1 + i)
					val := (duration % (time.Duration(1+i) * time.Second)).Seconds()
					y := 100 + i*30
					return spot.List(
						&cocoa.ProgressIndicator{
							X: 10, Y: y, Width: 90, Height: 25,
							Min: 0, Max: max, Value: val,
						},
						&cocoa.Slider{
							X: 110, Y: y, Width: 25, Height: 25,
							Min: 0, Max: max, Value: val,
							Type: gocoa.SliderTypeCircular,
							OnValueChanged: func(value float64) {
								setCounter(int(value))
							},
						},
						&cocoa.TextField{
							X: 140, Y: y, Width: 50, Height: 25,
							Editable: true, Selectable: true, Bezeled: true,
							Value: fmt.Sprintf("%.0f%%", val/max*100),
						})
				}),
				&cocoa.ComboBox{
					X: 210, Y: 220, Width: 180, Height: 25,
					Items:         []string{"Zero", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"},
					SelectedIndex: counter % 10,
				},
				&cocoa.Slider{
					X: 210, Y: 190, Width: 180, Height: 25,
					Min: 0, Max: 9, Value: float64(counter % 10),
					OnValueChanged: func(value float64) {
						setCounter(int(value))
					},
					Type: gocoa.SliderTypeLinear,
				},
				&cocoa.Button{
					X:      210,
					Y:      160,
					Width:  180,
					Height: 25,
					Title:  buttonTitle,
					OnClick: func() {
						setCounter(counter + 1)
					},
				},
				// Make2(QuitButton),
				ctx.Make(QuitButton),
				&cocoa.TextField{
					X: 10, Y: 10, Width: 380, Height: 80,
					Value:    RandStringBytesMaskImprSrcSB(300),
					Editable: false, Selectable: false, Bezeled: false, NoBackground: false,
				},
			},
		}
	}).Mount() //.Build()

	// title, setTitle := UseState[string]("Hello world!")

	// wnd := &Window{
	// 	Title:  title,
	// 	Width:  500,
	// 	Height: 500,
	// 	Children: []any{
	// 		&Button{
	// 			X:      LeftMargin,
	// 			Y:      BottomMargin + 255 + WidgetSpacing,
	// 			Width:  160,
	// 			Height: ButtonHeight,
	// 			Title:  "Generate Diff",
	// 			OnClick: func() {
	// 				fmt.Println("Diff generated!")
	// 				setTitle("Diff Generated!")
	// 			},
	// 		},
	// 		&Button{
	// 			X:      75,
	// 			Y:      50,
	// 			Width:  150,
	// 			Height: 25,
	// 			Title:  "Quit",
	// 			OnClick: func() {
	// 				gocoa.TerminateApplication()
	// 			},
	// 		},
	// 	},
	// }
	// wnd.Build()

	// wnd := gocoa.NewWindow("Hello World!", 150, 150, LeftMargin+160+WidgetSpacing+550+RightMargin, BottomMargin+255+WidgetSpacing+CommitListHeight+TopMargin)
	// width, _ := wnd.Size()

	// fileList := gocoa.NewTextView(LeftMargin, BottomMargin+255+WidgetSpacing+30+WidgetSpacing, FileListWidth, 100)
	// fileList.SetText("File list here\nBlablab")
	// wnd.AddTextView(fileList)

	// commitList := gocoa.NewTextView(LeftMargin+160+WidgetSpacing, BottomMargin+255+WidgetSpacing, width-LeftMargin-RightMargin-WidgetSpacing-FileListWidth, CommitListHeight)
	// commitList.SetText("Commit list here\nBlablab")
	// wnd.AddTextView(commitList)

	// // Change me button
	// currentTitle, nextTitle := "Generate Diff", "Diff Generated!"
	// changeButton := gocoa.NewButton(LeftMargin, BottomMargin+255+WidgetSpacing, 160, ButtonHeight)
	// changeButton.SetTitle(currentTitle)
	// changeButton.SetBezelStyle(gocoa.ButtonBezelStyleRoundRect)
	// changeButton.OnClick(func() {
	// 	changeButton.SetTitle(nextTitle)
	// 	currentTitle, nextTitle = nextTitle, currentTitle
	// })
	// wnd.AddButton(changeButton)

	// diffView := gocoa.NewTextView(LeftMargin, BottomMargin, 160+WidgetSpacing+550, 255)
	// diffView.SetText("Diff view here\nBlablab")
	// wnd.AddTextView(diffView)

	// // Quit button
	// quitButton := gocoa.NewButton(75, 50, 150, 25)
	// quitButton.SetTitle("Quit")
	// quitButton.OnClick(func() { gocoa.TerminateApplication() })
	// wnd.AddButton(quitButton)

	// wnd.MakeKeyAndOrderFront()
	// wnd.OnDidMove(func(w *gocoa.Window) {
	// 	fmt.Println("Window moved!")
	// })
	// wnd.OnDidResize(func(w *gocoa.Window) {
	// 	width, height := w.Size()
	// 	fmt.Printf("Window resized to %dx%d\n", width, height)
	// })

	gocoa.RunApplication()
}
