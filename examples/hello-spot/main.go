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

type BlinkingLabel struct {
	X, Y, Width, Height int
	Text                string
	Size                int
}

func (b *BlinkingLabel) Render(ctx *spot.RenderContext) spot.Component {
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

	txt := b.Text
	if !visible {
		txt = "🙈"
	}

	return &ui.Label{
		X: b.X, Y: b.Y, Width: b.Width, Height: b.Height,
		Value:    txt,
		FontSize: b.Size,
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
		X:      210,
		Y:      370,
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
	spot.MountFn(func(ctx *spot.RenderContext) spot.Component {
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
		randText := RandStringBytesMaskImprSrcSB(400)
		if duration-duration.Truncate(time.Second) < time.Millisecond*200 {
			randText = strings.Repeat(" ", 20) + "▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄▄\n" +
				strings.Repeat(" ", 20) + "██░▄▄▄░█▀▄▄▀█▀▄▄▀█▄░▄█░██\n" +
				strings.Repeat(" ", 20) + "██▄▄▄▀▀█░▀▀░█░██░██░██▄██\n" +
				strings.Repeat(" ", 20) + "██░▀▀▀░█░█████▄▄███▄██▀██\n" +
				strings.Repeat(" ", 20) + "▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀▀"
		}

		return &ui.Window{
			Title: "Hello Spot!",
			Width: 400, Height: 500,
			Children: []spot.Component{
				&ui.Label{
					X: 10, Y: 410, Width: 380, Height: 80,
					Value: fmt.Sprintf("%02d:%02d.%03d", int(duration.Minutes())%60,
						int(duration.Seconds())%60, duration.Milliseconds()%1000),
					FontSize: 60,
					Align:    ui.LabelAlignmentCenter,
					// Editable: false, Selectable: false, Bezeled: false, NoBackground: false,
				},
				spot.Range(ctx, 0, 10, func(ctx *spot.RenderContext, i int) spot.Component {
					max := float64(1 + i)
					val := (duration % (time.Duration(1+i) * time.Second)).Seconds()
					y := 100 + i*30

					return spot.Fragment{
						&ui.ProgressBar{
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
					}
				}),
				&ui.Dropdown{
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
				&BlinkingLabel{X: 210, Y: 120, Width: 180, Height: 30, Text: ui.BackendName, Size: 20},
				&ui.ListBox{
					X: 210, Y: 285, Width: 180, Height: 75,
					Values:    []string{"Null", "Eins", "Zwei", "Drei", "Vier", "Fünf", "Sechs", "Sieben", "Acht", "Neun"},
					Selection: []int{counter % 10},
					OnSelect: func(selection []int) {
						if len(selection) > 0 {
							setCounter(selection[0])
						}
					},
				},
				spot.Make(QuitButton),
				&ui.TextView{
					X: 10, Y: 10, Width: 380, Height: 80,
					Text: randText,
					// FontSize: 11,
					// Editable: false, Selectable: false, Bezeled: false, NoBackground: false,
				},
			},
		}
	})

	ui.Run()
}
