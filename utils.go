package selo

func wrapFactory[T any](f Factory[T]) Factory[any] {
	return func() any {
		return f()
	}
}
