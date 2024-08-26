package selo

type uniqueAccessor struct {
	f        Factory[any]
	v        any
	hasValue bool
}

func (sa *uniqueAccessor) Get() any {
	if sa.hasValue == false {
		sa.v = sa.f()
	}
	return sa.v
}

type uniqueAccessorBuilder[T any] struct {
	commonAccessorSettings
}

func (b *uniqueAccessorBuilder[T]) SetTag(tag string) AccessorBuilder[T] {
	b.tag = tag
	return b
}

func (b *uniqueAccessorBuilder[T]) SetFactory(f Factory[T]) AccessorBuilder[T] {
	b.f = wrapFactory(f)
	return b
}

func (b *uniqueAccessorBuilder[T]) SetLazy(lazy bool) AccessorBuilder[T] {
	b.lazy = lazy
	return b
}

func (b *uniqueAccessorBuilder[T]) Record() {
	a := new(uniqueAccessor)
	a.f = b.f
	a.hasValue = !b.lazy

	if b.lazy == false {
		a.v = b.f()
	}

	l.set((*T)(nil), a)
}
