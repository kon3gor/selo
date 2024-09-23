package benchmarks

import "testing"

func BenchmarkIntKeys(b *testing.B) {
	store := NewStore[int]()
	b.Run("set", func(b *testing.B) {
		for i := range b.N {
			store.Put(i)
		}
	})

	b.Run("get", func(b *testing.B) {
		for i := range b.N {
			_ = store.Get(i)
		}
	})
}
