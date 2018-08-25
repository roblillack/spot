package main

import (
	"github.com/mojbro/gocoa"
)

func main() {
	app := gocoa.GetSharedApplication()
	wnd := gocoa.NewWindow("Hello World!", 150, 150, 600, 400)
	button := gocoa.NewButton()
	button.SetTitle("Quit")
	wnd.AddButton(button)
	wnd.MakeKeyAndOrderFront()
	app.Run()
}
