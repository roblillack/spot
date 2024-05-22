package main

import (
	"fmt"
	"time"

	"github.com/mojbro/gocoa"
)

var wnd *gocoa.Window
var indicator *gocoa.ProgressIndicator

const maxValue = 100.00
const windowTitle = "Hello World!"

func main() {
	gocoa.InitApplication()
	gocoa.OnApplicationDidFinishLaunching(func() {
		fmt.Println("App running!")
	})
	wnd = gocoa.NewCenteredWindow(windowTitle, 320, 240)
	wnd.OnDidMove(func(uwnd *gocoa.Window) {
		fmt.Printf("old: %v\nnew: %v\n", wnd, uwnd)
	})

	// TextField
	titleTextField := gocoa.NewTextField(85, 195, 90, 20)
	titleTextField.SetStringValue(windowTitle)

	wnd.AddTextField(titleTextField)
	setTitleButton := gocoa.NewButton(175, 180, 60, 50)
	setTitleButton.SetTitle("SET")
	setTitleButton.OnClick(func() {
		wnd.SetTitle(titleTextField.StringValue())
	})
	wnd.AddButton(setTitleButton)

	// TextView
	textView := gocoa.NewTextView(85, 165, 90, 20)
	textView.SetText("lorem ipsum")
	remTextButton := gocoa.NewButton(175, 150, 60, 50)
	remTextButton.SetTitle("ADD")
	visible := false
	remTextButton.OnClick(func() {
		if visible {
			remTextButton.SetTitle("ADD")
			textView.Remove()
			visible = false
		} else {
			remTextButton.SetTitle("DEL")
			wnd.AddTextView(textView)
			visible = true
		}
	})
	wnd.AddButton(remTextButton)

	// Change me button
	currentTitle, nextTitle := "Change me!", "Change me again!"
	changeButton := gocoa.NewButton(75, 120, 160, 25)
	changeButton.SetTitle(currentTitle)
	changeButton.OnClick(func() {
		changeButton.SetTitle(nextTitle)
		currentTitle, nextTitle = nextTitle, currentTitle
	})
	wnd.AddButton(changeButton)

	// Download button
	loadButton := gocoa.NewButton(75, 95, 160, 25)
	loadButton.SetTitle("download")

	loadButton.OnClick(func() {
		indicator.Show()
		go func() {
			increase()
		}()
	})
	wnd.AddButton(loadButton)

	// ProgressIndicator
	indicator = gocoa.NewProgressIndicator(75, 75, 160, 25)
	indicator.Hide()

	indicator.SetLimits(0.00, maxValue)
	indicator.SetValue(0.00)
	indicator.SetAutohide(true)
	wnd.AddProgressIndicator(indicator)

	// Quit button
	quitButton := gocoa.NewButton(75, 45, 160, 25)
	quitButton.SetTitle("Quit")
	quitButton.OnClick(func() {
		gocoa.TerminateApplication()
	})
	wnd.AddButton(quitButton)

	infoView := gocoa.NewTextView(80, 15, 150, 5)
	infoView.SetText(fmt.Sprintf("screen size %v\n", wnd.GetScreen()))
	wnd.AddTextView(infoView)

	wnd.MakeKeyAndOrderFront()
	gocoa.RunApplication()

}

func increase() {
	time.Sleep(1 * time.Second)
	value := indicator.GetValue()
	inc := 17.00
	indicator.IncrementBy(inc)
	fmt.Printf("Increasing… %f by %f\n", value, inc)
	if indicator.GetValue() < maxValue {
		increase()
	}
}
