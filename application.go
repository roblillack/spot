package gocoa

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "application.h"
import "C"
import "unsafe"

// Application object
type Application struct {
	sharedApp unsafe.Pointer
}

// GetSharedApplication Returns an application object pointing to the Cocoa shared application object.
func GetSharedApplication() *Application {
	return &Application{sharedApp: C.GetSharedApplication()}
}

// Run starts the main runloop
func (app *Application) Run() {
	C.App_Run(app.sharedApp)
}
