package ui

import (
	"fmt"
	"testing"

	"github.com/roblillack/spot"
)

func BenchmarkMeaninglessUpdates(b *testing.B) {
	node := spot.Build(&Button{})
	node.Mount()

	for i := 0; i < b.N; i++ {
		node.Update(spot.Build(&Button{}), nil)
	}
}

func BenchmarkSimpleUpdates(b *testing.B) {
	node := spot.Build(&Button{})
	node.Mount()

	for i := 0; i < b.N; i++ {
		node.Update(spot.Build(&Button{Title: fmt.Sprintf("%d", i)}), nil)
	}
}
