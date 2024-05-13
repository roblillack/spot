package main

import (
	"fmt"
	"journey/fltk"
	"journey/spot"
	"math/rand"
	"os"
	"strings"
	"time"

	goFltk "github.com/pwiecz/go-fltk"
)

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
		return &fltk.Button{
			X:      210,
			Y:      100,
			Width:  180,
			Height: 25,
			Title:  "Please wait...",
		}
	}

	return &fltk.Button{
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

		return &fltk.Window{
			Title: "Hello World!",
			Width: 400, Height: 500,
			Children: []spot.Component{
				&fltk.Label{
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
						&fltk.ProgressIndicator{
							X: 10, Y: y, Width: 90, Height: 25,
							Min: 0, Max: max, Value: val,
						},
						&fltk.Slider{
							X: 110, Y: y, Width: 25, Height: 25,
							Min: 0, Max: max, Value: val,
							// Type: gocoa.SliderTypeCircular,
							OnValueChanged: func(value float64) {
								setCounter(int(value))
							},
						},
						&fltk.TextField{
							X: 140, Y: y, Width: 50, Height: 25,
							// Editable: true, Selectable: true, Bezeled: true,
							Value: fmt.Sprintf("%.0f%%", val/max*100),
						},
					)
				}),
				&fltk.ComboBox{
					X: 210, Y: 220, Width: 180, Height: 25,
					Items:         []string{"Zero", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"},
					SelectedIndex: counter % 10,
					OnSelectionDidChange: func(index int) {
						setCounter(index)
					},
				},
				&fltk.Spinner{
					X: 210, Y: 250, Width: 50, Height: 25,
					Min: 0, Max: 9, Step: 1, Value: float64(counter % 10),
					OnValueChanged: func(value float64) {
						setCounter(int(value))
					},
				},
				&fltk.Checkbox{
					X: 270, Y: 250, Width: 120, Height: 25,
					Checked: counter%2 == 0,
					Label:   "is even?",
					OnChange: func(checked bool) {
						setCounter(counter + 1)
					},
				},
				&fltk.Slider{
					X: 210, Y: 190, Width: 180, Height: 25,
					Min: 0, Max: 9, Value: float64(counter % 10),
					OnValueChanged: func(value float64) {
						setCounter(int(value))
					},
					// Type: gocoa.SliderTypeLinear,
				},
				&fltk.Button{
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
				&fltk.TextField{
					X: 10, Y: 10, Width: 380, Height: 80,
					Value: RandStringBytesMaskImprSrcSB(300),
					// Editable: false, Selectable: false, Bezeled: false, NoBackground: false,
				},
			},
		}
	}).Mount()

	goFltk.Run()
}
