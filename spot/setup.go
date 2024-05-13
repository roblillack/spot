package spot

var RunOnMainLoop func(func()) = func(_ func()) {
	panic("RunOnMainLoop not set up.")
}
