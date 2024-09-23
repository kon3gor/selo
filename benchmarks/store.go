package benchmarks

type Store[K comparable] struct {
	i map[K]any
}

func (s *Store[K]) Put(k K) {
	s.i[k] = struct{}{}
}

func (s *Store[K]) Get(k K) any {
	return s.i[k]
}

func NewStore[K comparable]() *Store[K] {
	s := new(Store[K])
	s.i = make(map[K]any)
	return s
}
