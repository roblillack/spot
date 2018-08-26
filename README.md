# gocoa
[![GoDoc](https://godoc.org/github.com/mojbro/gocoa?status.svg)](https://godoc.org/github.com/mojbro/gocoa)
[![Go Report Card](https://goreportcard.com/badge/github.com/mojbro/gocoa)](https://goreportcard.com/report/github.com/mojbro/gocoa)

Go bindings for the Cocoa framework to build macOS applications.

<img src="resources/images/helloworld-screenshot.png" width="412" />

## How to use

The following is a basic [Hello World](examples/helloworld) application.

```go
package main

import (
	"fmt"

	"github.com/mojbro/gocoa"
)

func main() {
	gocoa.InitApplication()
	gocoa.OnApplicationDidFinishLaunching(func() {
		fmt.Println("App running!")
	})
	wnd := gocoa.NewWindow("Hello World!", 150, 150, 300, 200)

	// Change me button
	currentTitle, nextTitle := "Change me!", "Change me again!"
	changeButton := gocoa.NewButton(75, 125, 150, 25)
	changeButton.SetTitle(currentTitle)
	changeButton.OnClick(func() {
		changeButton.SetTitle(nextTitle)
		currentTitle, nextTitle = nextTitle, currentTitle
	})
	wnd.AddButton(changeButton)

	// Quit button
	quitButton := gocoa.NewButton(75, 50, 150, 25)
	quitButton.SetTitle("Quit")
	quitButton.OnClick(func() { gocoa.TerminateApplication() })
	wnd.AddButton(quitButton)

	wnd.MakeKeyAndOrderFront()
	gocoa.RunApplication()
}
```

## Status of this project

This package is very, very early and incomplete! It is mostly just an experiment and is not really
useful yet. I will continue working on it with the aim to be able to make useful native macOS apps in Go.
