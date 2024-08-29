package selo

type taggedKey[T any] struct {
	v   *T
	tag string
}

func newTaggedKey[T any](tag string) any {
	return taggedKey[T]{
		v:   (*T)(nil),
		tag: tag,
	}
}
