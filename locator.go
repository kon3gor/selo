package selo

type Factory[T any] func() T

func wrapFactory[T any](f Factory[T]) Factory[any] {
	return func() any {
		return f()
	}
}

type Accessor interface {
	Get() any
}

type Locator interface {
	set(any, Accessor)
}

type locator struct {
	accessors map[any]Accessor
}

var _ Locator = &locator{}

var l Locator

func Init(opts ...Option) {
	l = &locator{
		accessors: make(map[any]Accessor),
	}
}

func (l *locator) set(k any, a Accessor) {
	l.accessors[k] = a
}

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

func Get[T any]() T {
	return l.(*locator).accessors[(*T)(nil)].Get().(T)
}

func Unique[T any]() AccessorBuilder[T] {
	//todo: use sync.Pool for builders
	return &uniqueAccessorBuilder[T]{}
}
