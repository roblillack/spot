package main

import (
	"github.com/mojbro/gocoa"
)

func main() {
	gocoa.InitApplication()
	width, height := 300, 200
	wnd := gocoa.NewWindow("Hello World!", 150, 150, width, height)
	buttonWidth, buttonHeight := 150, 24
	button := gocoa.NewButton(width/2-buttonWidth/2, height/2-buttonHeight/2, buttonWidth, buttonHeight)
	button.SetTitle("Click to quit app!")
	wnd.AddButton(button)
	wnd.MakeKeyAndOrderFront()
	gocoa.RunApplication()
}
