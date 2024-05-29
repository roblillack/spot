package cocoa

import (
	"fmt"
	"testing"
)

func TestNewProgressIndicator(t *testing.T) {
	InitApplication()
	OnApplicationDidFinishLaunching(func() {
		fmt.Println("App running!")
	})
	wnd := NewWindow("Hello World!", 150, 150, 300, 200)

	indicator := NewProgressIndicator(0, 0, 100, 50)
	if indicator.progressIndicatorPtr == nil {
		t.Fatalf("pointer to C indicator is nil!")
	}

	if indicator.GetValue() != 0.00 {
		t.Fatalf("indicator has not set the default value, but has value %f!", indicator.GetValue())
	}

	wnd.AddProgressIndicator(indicator)
	indicator.SetValue(66.00)
	indicator.IncrementBy(5.00)
	indicator.IncrementBy(-20.00)

	TerminateApplication()
}

func TestSetValue(t *testing.T) {
	indicator := NewProgressIndicator(0, 0, 100, 50)
	indicator.SetLimits(0.00, 100.00)

	indicator.SetValue(33.00)
	if indicator.GetValue() != 33.00 {
		t.Fatalf("indicator did not update the value, but has value %f!", indicator.GetValue())
	}
}

func TestIncrementBy(t *testing.T) {
	indicator := NewProgressIndicator(0, 0, 100, 50)
	indicator.SetLimits(0.00, 100.00)

	indicator.IncrementBy(1.00)
	indicator.IncrementBy(1.00)
	indicator.IncrementBy(1.00)
	indicator.IncrementBy(1.00)
	indicator.IncrementBy(1.00)
	if indicator.GetValue() != 5.00 {
		t.Fatalf("indicator did not increment the value but %f!", indicator.GetValue())
	}

	indicator.IncrementBy(-1.00)
	if indicator.GetValue() != 4.00 {
		t.Fatalf("indicator did not decrement the value but %f!", indicator.GetValue())
	}
}
