package selo

type Locator interface {
	set(int, anyAccessor)
}

type locator struct {
	accessors map[int]Accessor[any]
}

var _ Locator = &locator{}

var l Locator

func Init(opts ...Option) {
	l = &locator{
		accessors: make(map[int]Accessor[any]),
	}
}

func (l *locator) set(k int, a anyAccessor) {
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

	l.set(key, newSingletonAccessor(f))
}

func Get[T any](key int) T {
	return l.(*locator).accessors[key].Get().(T)
}
