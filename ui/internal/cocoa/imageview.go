package cocoa

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #import "imageview.h"
// #import "view.h"
// #include <stdlib.h>
import "C"
import (
	"fmt"
	"image"
)

// Represents an ImageView control that can display images.
type ImageView struct {
	ptr      C.ImageViewPtr
	callback func()
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

func NewImageViewWithContentsOfURL(x int, y int, width int, height int, url string) *ImageView {
	imageViewID := len(imageViews)
	imageViewPtr := C.ImageView_NewWithContentsOfURL(C.int(imageViewID), C.int(x), C.int(y), C.int(width), C.int(height), C.CString(url))

	img := &ImageView{
		ptr: imageViewPtr,
	}
	imageViews = append(imageViews, img)
	return img
}

func NewImageView(x int, y int, width int, height int) *ImageView {
	imageViewID := len(imageViews)
	imageViewPtr := C.ImageView_New(C.int(imageViewID), C.int(x), C.int(y), C.int(width), C.int(height))

	img := &ImageView{
		ptr: imageViewPtr,
	}
	imageViews = append(imageViews, img)

	return img
}

func NewImageViewWithImage(x int, y int, width int, height int, image *image.RGBA) *ImageView {
	img := NewImageView(x, y, width, height)
	img.SetImage(image)
	return img
}

func (imageView *ImageView) SetImage(img *image.RGBA) {
	bytes := C.CBytes(img.Pix)
	nsImage := C.Image_NewWithRGBA(C.int(img.Bounds().Dx()), C.int(img.Bounds().Dy()), (*C.uchar)(bytes))
	C.ImageView_SetImage(imageView.ptr, nsImage)
	C.free(bytes)
}

func (imageView *ImageView) SetEditable(editable bool) {
	if editable {
		C.ImageView_SetEditable(imageView.ptr, 1)
	} else {
		C.ImageView_SetEditable(imageView.ptr, 0)
	}
}

func (imageView *ImageView) SetImageFrameStyle(frameStyle FrameStyle) {
	C.ImageView_SetFrameStyle(imageView.ptr, C.int(frameStyle))
}

func (imageView *ImageView) SetImageAlignment(imageAlignment ImageAlignment) {
	C.ImageView_SetImageAlignment(imageView.ptr, C.int(imageAlignment))
}

func (imageView *ImageView) SetImageScaling(imageScaling ImageScaling) {
	C.ImageView_SetImageScaling(imageView.ptr, C.int(imageScaling))
}

func (imageView *ImageView) SetAnimates(animates bool) {
	if animates {
		C.ImageView_SetAnimates(imageView.ptr, 1)
	} else {
		C.ImageView_SetAnimates(imageView.ptr, 0)
	}
}

func (imageView *ImageView) SetContentTintColor(hexRGBA string) {
	var r, g, b, a = 0, 0, 0, 0
	fmt.Sscanf(hexRGBA, "#%02x%02x%02x%02x", &r, &g, &b, &a)
	C.ImageView_SetContentTintColor(imageView.ptr, C.int(r), C.int(g), C.int(b), C.int(a))
}

// Remove removes an ImageView from the parent view again.
func (imageView *ImageView) Remove() {
	C.ImageView_Remove(imageView.ptr)
}

func (c *ImageView) SetFrameOrigin(x, y int) {
	C.View_SetFrameOrigin(C.ViewPtr(c.ptr), C.int(x), C.int(y))
}

func (c *ImageView) SetFrameSize(width, height int) {
	C.View_SetFrameSize(C.ViewPtr(c.ptr), C.int(width), C.int(height))
}

func (c *ImageView) SetFrame(x, y, width, height int) {
	C.View_SetFrame(C.ViewPtr(c.ptr), C.int(x), C.int(y), C.int(width), C.int(height))
}

func (c *ImageView) Frame() (x, y, width, height int) {
	var x_, y_, width_, height_ C.int
	C.View_Frame(C.ViewPtr(c.ptr), &x_, &y_, &width_, &height_)
	return int(x_), int(y_), int(width_), int(height_)
}
