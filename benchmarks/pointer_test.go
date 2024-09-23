package benchmarks

import "testing"

func BenchmarkPointerKeys(b *testing.B) {
	type Key struct {
		buffer [256]byte
	}
	store := NewStore[any]()
	b.Run("set", func(b *testing.B) {
		for range b.N {
			store.Put((*Key)(nil))
		}
	})

	b.Run("get", func(b *testing.B) {
		for range b.N {
			_ = store.Get((*Key)(nil))
		}
	})
}
