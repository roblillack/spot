package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/roblillack/spot"
	"github.com/roblillack/spot/ui"
)

func BlinkingLabel(ctx *spot.RenderContext, x, y, width, height int, text string, size int) spot.Component {
	visible, setVisible := spot.UseState(ctx, true)
	spot.UseEffect(ctx, func() {
		go func() {
			val := visible
			for {
				time.Sleep(500 * time.Millisecond)
				val = !val
				setVisible(val)
			}
		}()
	}, []any{})

	txt := text
	if !visible {
		txt = "🙈"
	}

	return &ui.Label{
		X: x, Y: y, Width: width,
		Height: height, Value: txt,
		FontSize: size,
	}
}

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
		return &ui.Button{
			X:      210,
			Y:      370,
			Width:  180,
			Height: 25,
			Title:  "Please wait...",
		}
	}

	return &ui.Button{
		X:      200,
		Y:      100,
		Width:  180,
		Height: 25,
		Title:  "Quit",
		OnClick: func() {
			os.Exit(0)
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
	ui.Init()

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

		return &ui.Window{
			Title: "Hello World!",
			Width: 400, Height: 500,
			Children: []spot.Component{
				&ui.Label{
					X: 10, Y: 410, Width: 380, Height: 80,
					Value: fmt.Sprintf("%02d:%02d.%03d", int(duration.Minutes())%60,
						int(duration.Seconds())%60, duration.Milliseconds()%1000),
					FontSize: 60,
					// Editable: false, Selectable: false, Bezeled: false, NoBackground: false,
				},
				spot.Range(ctx, 0, 10, func(ctx *spot.RenderContext, i int) spot.Component {
					max := float64(1 + i)
					val := (duration % (time.Duration(1+i) * time.Second)).Seconds()
					y := 100 + i*30
					return spot.List(
						&ui.ProgressIndicator{
							X: 10, Y: y, Width: 90, Height: 25,
							Min: 0, Max: max, Value: val,
						},
						&ui.Dial{
							X: 110, Y: y, Width: 25, Height: 25,
							Min: 0, Max: max, Value: val,
							OnValueChanged: func(value float64) {
								setCounter(int(value))
							},
						},
						&ui.TextField{
							X: 140, Y: y, Width: 50, Height: 25,
							// Editable: true, Selectable: true, Bezeled: true,
							Value: fmt.Sprintf("%.0f%%", val/max*100),
						},
					)
				}),
				&ui.ComboBox{
					X: 210, Y: 220, Width: 180, Height: 25,
					Items:         []string{"Zero", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"},
					SelectedIndex: counter % 10,
					OnSelectionDidChange: func(index int) {
						setCounter(index)
					},
				},
				&ui.Spinner{
					X: 210, Y: 250, Width: 50, Height: 25,
					Min: 0, Max: 9, Step: 1, Value: float64(counter % 10),
					OnValueChanged: func(value float64) {
						setCounter(int(value))
					},
				},
				&ui.Checkbox{
					X: 270, Y: 250, Width: 120, Height: 25,
					Checked: counter%2 == 0,
					Label:   "is even?",
					OnChange: func(checked bool) {
						setCounter(counter + 1)
					},
				},
				&ui.Slider{
					X: 210, Y: 190, Width: 180, Height: 25,
					Min: 0, Max: 9, Value: float64(counter % 10),
					OnValueChanged: func(value float64) {
						setCounter(int(value))
					},
					// Type: gocoa.SliderTypeLinear,
				},
				&ui.Button{
					X:      210,
					Y:      160,
					Width:  180,
					Height: 25,
					Title:  buttonTitle,
					OnClick: func() {
						setCounter(counter + 1)
					},
				},
				&ui.Label{X: 210, Y: 100, Width: 180, Height: 25, Value: "Current backend:"},
				ctx.Make(func(x *spot.RenderContext) spot.Component {
					return BlinkingLabel(x, 210, 120, 180, 30, ui.BackendName, 20)
				}),
				ctx.Make(QuitButton),
				&ui.TextField{
					X: 10, Y: 10, Width: 380, Height: 80,
					Value: RandStringBytesMaskImprSrcSB(300),
					// Editable: false, Selectable: false, Bezeled: false, NoBackground: false,
				},
			},
		}
	}).Mount()

	ui.Run()
}
