package selo

type Locator interface {
	set(any, Accessor)
}

type locator struct {
	accessors map[any]Accessor
}

var _ Locator = &locator{}

var l Locator

func Init(opts ...Option) {
	//todo: make it actualy matter
	s := new(locatorSettings)
	for _, opt := range opts {
		opt(s)
	}

	l = &locator{
		accessors: make(map[any]Accessor),
	}
}

func (l *locator) set(k any, a Accessor) {
	l.accessors[k] = a
}

func Get[T any]() T {
	return l.(*locator).accessors[(*T)(nil)].Get().(T)
}

func Unique[T any]() AccessorBuilder[T] {
	//todo: use sync.Pool for builders
	return &uniqueAccessorBuilder[T]{}
}
