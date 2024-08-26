package selo

type Accessor interface {
	Get() any
}

type Factory[T any] func() T

type AccessorBuilder[T any] interface {
	SetTag(tag string) AccessorBuilder[T]
	SetFactory(f Factory[T]) AccessorBuilder[T]
	SetLazy(lazy bool) AccessorBuilder[T]
	Record()
}

type commonAccessorSettings struct {
	tag  string
	lazy bool
	f    Factory[any]
}
