package gocoa

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #import "imageview.h"
import "C"

// Represents an ImageView control that can display images.
type ImageView struct {
	imageViewPtr C.ImageViewPtr
	callback     func()
}

type FrameStyle int32
type ImageAlignment int32
type ImageScaling int32

const (
	FrameStyleNone      FrameStyle = 0
	FrameStylePhoto     FrameStyle = 1
	FrameStyleGrayBezel FrameStyle = 2
	FrameStyleGroove    FrameStyle = 3
	FrameStyleButton    FrameStyle = 4
)

const (
	ImageAlignCenter      ImageAlignment = 0
	ImageAlignTop         ImageAlignment = 1
	ImageAlignTopLeft     ImageAlignment = 2
	ImageAlignTopRight    ImageAlignment = 3
	ImageAlignLeft        ImageAlignment = 4
	ImageAlignBottom      ImageAlignment = 5
	ImageAlignBottomLeft  ImageAlignment = 6
	ImageAlignBottomRight ImageAlignment = 7
	ImageAlignRight       ImageAlignment = 8
)

const (
	ImageScalingScaleProportionallyDown     ImageScaling = 0
	ImageScalingScaleAxesIndependently      ImageScaling = 1
	ImageScalingScaleNone                   ImageScaling = 2
	ImageScalingScaleProportionallyUpOrDown ImageScaling = 3
)

var imageViews []*ImageView

func NewImageView(x int, y int, width int, height int, url string) *ImageView {
	imageViewID := len(imageViews)
	imageViewPtr := C.ImageView_New(C.int(imageViewID), C.int(x), C.int(y), C.int(width), C.int(height), C.CString(url))

	img := &ImageView{
		imageViewPtr: imageViewPtr,
	}
	imageViews = append(imageViews, img)
	return img
}

func (imageView *ImageView) SetEditable(editable bool) {

}

func (imageView *ImageView) SetAllowsCutCopyPaste(cutCopyPaste bool) {

}

func (imageView *ImageView) SetImageFrameStyle(frameStyle FrameStyle) {
	C.ImageView_SetFrameStyle(imageView.imageViewPtr, C.int(frameStyle))
}

func (imageView *ImageView) SetImageAlignment(imageAlignment ImageAlignment) {
	C.ImageView_SetImageAlignment(imageView.imageViewPtr, C.int(imageAlignment))
}

func (imageView *ImageView) SetImageScaling(imageScaling ImageScaling) {
	C.ImageView_SetImageScaling(imageView.imageViewPtr, C.int(imageScaling))
}

func (imageView *ImageView) SetAnimates(animates bool) {

}
