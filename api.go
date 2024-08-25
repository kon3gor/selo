package selo

type Factory[T any] func() T

type AnyFactory Factory[any]

func wrapFactory[T any](f Factory[T]) Factory[any] {
	return func() any {
		return f()
	}
}

type Accessor[T any] interface {
	Get() T
}

type anyAccessor Accessor[any]

type singletonAccessor[T any] struct {
	f        Factory[T]
	v        T
	hasValue bool
}

var _ Accessor[int] = &singletonAccessor[int]{}

func newSingletonAccessor[T any](f Factory[T]) anyAccessor {
	wf := wrapFactory(f)

	return &singletonAccessor[any]{
		f:        wf,
		v:        wf(),
		hasValue: true,
	}
}

func (sa *singletonAccessor[T]) Get() T {
	if sa.hasValue == false {
		sa.v = sa.f()
	}

	return sa.v
}
