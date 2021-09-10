package main

import (
	"os"
	"path/filepath"

	"github.com/mojbro/gocoa"
)

func main() {
	gocoa.InitApplication()

	wnd := gocoa.NewCenteredWindow("ImageView example", 400, 300)

	// Image 1: Remote URL
	imageView := gocoa.NewImageView(75, 175, 200, 100, "https://www.google.com/images/branding/googlelogo/1x/googlelogo_color_272x92dp.png")
	imageView.SetImageFrameStyle(gocoa.FrameStyleNone)
	imageView.SetImageAlignment(gocoa.ImageAlignBottomRight)
	imageView.SetImageScaling(gocoa.ImageScalingScaleProportionallyUpOrDown)
	imageView.SetContentTintColor("#FF0000FF")
	imageView.SetEditable(false)
	wnd.AddImageView(imageView)

	// get path of this executable
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	spyroPath := filepath.Join("file://", exPath, "/spyro.png") // file://<absolute_path>/spyro.png

	// Image 2: Local URL
	imageView2 := gocoa.NewImageView(175, 25, 100, 100, spyroPath)
	imageView2.SetImageFrameStyle(gocoa.FrameStyleNone)
	imageView2.SetImageScaling(gocoa.ImageScalingScaleProportionallyUpOrDown)
	imageView2.SetContentTintColor("#FF0000FF")
	wnd.AddImageView(imageView2)

	// Quit button
	quitButton := gocoa.NewButton(25, 50, 100, 25)
	quitButton.SetTitle("Quit")
	quitButton.OnClick(func() { gocoa.TerminateApplication() })
	wnd.AddButton(quitButton)

	wnd.MakeKeyAndOrderFront()
	gocoa.RunApplication()
}
