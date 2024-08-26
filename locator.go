package selo

type Locator interface {
	set(any, anyAccessor)
}

type locator struct {
	accessors map[any]Accessor[any]
}

var _ Locator = &locator{}

var l Locator

func Init(opts ...Option) {
	l = &locator{
		accessors: make(map[any]Accessor[any]),
	}
}

func (l *locator) set(k any, a anyAccessor) {
	l.accessors[k] = a
}

type uniqeSettings struct {
	lazy bool
}

type UniqeOption func(*uniqeSettings)

func WithLazy(lazy bool) UniqeOption {
	return func(us *uniqeSettings) {
		us.lazy = lazy
	}
}

func Unique[T any](key int, f Factory[T], opts ...UniqeOption) {
	s := new(uniqeSettings)
	for _, opt := range opts {
		opt(s)
	}

	l.set((*T)(nil), newSingletonAccessor(f))
}

func Get[T any](key int) T {
	return l.(*locator).accessors[(*T)(nil)].Get().(T)
}
