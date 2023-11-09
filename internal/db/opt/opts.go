package opt

type ctx struct {
	page     int
	pageSize int
	filter   map[string]any
}

type Opt func(ctx *ctx)

func WithPage(n int) Opt {
	return func(ctx *ctx) {
		ctx.page = n
	}
}

func WithPageSize(n int) Opt {
	return func(ctx *ctx) {
		ctx.pageSize = n
	}
}

func WithFilter(filter map[string]any) Opt {
	return func(ctx *ctx) {
		if ctx.filter == nil {
			ctx.filter = filter
			return
		}
		for k, v := range filter {
			ctx.filter[k] = v
		}
	}
}
