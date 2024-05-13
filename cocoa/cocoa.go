package cocoa

import (
	"journey/spot"

	"github.com/mojbro/gocoa"
)

func init() {
	spot.RunOnMainLoop = gocoa.RunOnMainLoop
}
