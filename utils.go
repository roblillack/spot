package spot

func Range[T float64 | int | uint](ctx *RenderContext, start, end T, f func(ctx *RenderContext, idx T) Component) Fragment {
	var res Fragment
	for i := start; i < end; i++ {
		res = append(res, f(ctx, i))
	}
	return res
}
