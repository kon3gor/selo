package selo

type Locator interface {
	set(any, Accessor)
	get(any) Accessor
}

type locator struct {
	accessors map[any]Accessor
	log       Logger
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
		log:       s.logger,
	}
}

func (l *locator) set(k any, a Accessor) {
	l.accessors[k] = a
}

func (l *locator) get(k any) Accessor {
	return l.accessors[k]
}

func Get[T any]() T {
	return l.get((*T)(nil)).Get().(T)
}

func GetTagged[T any](tag string) T {
	key := newTaggedKey[T](tag)
	return l.get(key).Get().(T)
}

func Unique[T any]() AccessorBuilder[T] {
	//todo: use sync.Pool for builders
	return &uniqueAccessorBuilder[T]{}
}
