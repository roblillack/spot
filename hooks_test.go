package spot

import (
	"testing"
)

func assertEq[T comparable](t *testing.T, expected T, actual T) {
	if expected != actual {
		t.Errorf("Expected %v != %v Actual", expected, actual)
	}
}

func assertNeq[T comparable](t *testing.T, expected T, actual T) {
	if expected == actual {
		t.Errorf("Unexpected equality: %v", actual)
	}
}

func TestUseState(t *testing.T) {
	ctx := &RenderContext{
		values: make(map[int]any),
	}

	ctx.count = 0
	v1, setStateV1 := UseState(ctx, 1)
	assertEq(t, 1, v1)
	ctx.count = 0
	v2, setStateV2 := UseState(ctx, 1)
	assertEq(t, 1, v2)
	setStateV1(11)
	ctx.count = 0
	v3, setStateV3 := UseState(ctx, 1)
	assertEq(t, 11, v3)
	setStateV3(33)
	setStateV2(22)
	ctx.count = 0
	v4, _ := UseState(ctx, 1)
	assertEq(t, 22, v4)
}

func TestUseEffect(t *testing.T) {
	ctx := &RenderContext{
		values: make(map[int]any),
	}

	counter := 0
	// Fake a render loop
	r := func(times int, deps []any) {
		for i := 0; i < times; i++ {
			ctx.count = 0
			UseEffect(ctx, func() { counter++ }, deps)
		}
	}

	// No deps should always run the effect
	r(7, nil)
	assertEq(t, 7, counter)

	// Same deps run the effect only once
	counter = 0
	ctx.values = make(map[int]any)
	r(7, []any{42, "test123"})
	assertEq(t, 1, counter)

	// Empty deps should run the effect only once
	counter = 0
	ctx.values = make(map[int]any)
	r(7, []any{})
	assertEq(t, 1, counter)

	// Changing deps should run the effect again
	counter = 0
	ctx.values = make(map[int]any)
	r(7, []any{42, "test123"})
	r(7, []any{23, "test123"})
	assertEq(t, 2, counter)
}

// func TestUseCallback(t *testing.T) {
// 	ctx := &RenderContext{
// 		values: make(map[int]any),
// 	}

// 	// Empty deps should always return the same callback
// 	var b *int
// 	var cb1, cb2 *Callback[func()]

// 	{
// 		a := 42
// 		cb1 = UseCallback(ctx, func() { val := a; b = &val }, []any{})
// 	}
// 	ctx.count = 0
// 	{
// 		a := 23
// 		cb2 = UseCallback(ctx, func() { val := a; b = &val }, []any{})
// 	}

// 	cb1.Fn()
// 	assertEq(t, 42, *b)
// 	cb2.Fn()
// 	assertEq(t, 42, *b)
// 	assertEq(t, cb1, cb2)

// 	// Changing deps should create a new callback
// 	ctx.values = make(map[int]any)
// 	ctx.count = 0
// 	{
// 		a := 42
// 		cb1 = UseCallback(ctx, func() { val := a; b = &val }, []any{a})
// 	}
// 	ctx.count = 0
// 	{
// 		a := 23
// 		cb2 = UseCallback(ctx, func() { val := a; b = &val }, []any{a})
// 	}

// 	cb1.Fn()
// 	assertEq(t, 42, *b)
// 	cb2.Fn()
// 	assertEq(t, 23, *b)
// 	assertNeq(t, cb1, cb2)
// }
