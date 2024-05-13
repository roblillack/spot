package spot

func UseState[T any](ctx *RenderContext, initial T) (T, func(next T)) {
	n := ctx.count
	ctx.count++

	setterFn := func(next T) {
		ctx.values[n] = next
		// ctx.changed = true
		ctx.TriggerUpdate()
	}
	if v, ok := ctx.values[n]; ok {
		// fmt.Printf("Using state %d for value: %v (initially %v)\n", n, v.(T), initial)
		return v.(T), setterFn
	}
	// fmt.Printf("Setting up state %d with initial value %v\n", n, initial)
	ctx.values[n] = initial
	return initial, setterFn
}

func UseEffect(ctx *RenderContext, fn func(), deps []any) {
	vals, setVals := UseState[[]any](ctx, nil)
	if deps == nil && vals != nil || (vals != nil && len(deps) != len(vals)) {
		panic("UseEffect: Length of dependencies changed")
	}

	changed := false
	if vals == nil {
		// first call: Always run the effect
		changed = true
	} else {
		for i, dep := range deps {
			if dep != vals[i] {
				changed = true
				break
			}
		}
	}

	if changed {
		setVals(deps)
		fn()
	}
}
