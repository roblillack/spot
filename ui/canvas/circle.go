package canvas

import (
	"image/color"
)

type Element interface{}

type Circle struct {
	X, Y        int
	Radius      int
	Fill        color.Color
	Stroke      color.Color
	StrokeWidth int
}
