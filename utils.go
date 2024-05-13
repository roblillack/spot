package spot

func Range[T float64 | int | uint](ctx *RenderContext, start, end T, f func(ctx *RenderContext, idx T) Component) ComponentList {
	var res ComponentList
	for i := start; i < end; i++ {
		res = append(res, f(ctx, i))
	}
	return res
}

func List(children ...Component) ComponentList {
	return children
}
