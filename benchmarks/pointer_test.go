package benchmarks

import "testing"

type VP = string
type KP = any
type MP = map[KP]VP

func BenchmarkPointerKeys(b *testing.B) {
	type SimpleStruct struct{}
	m := make(MP)
	b.Run("set", func(b *testing.B) {
		for range b.N {
			setPointer[SimpleStruct](m, "hello world")
		}
	})

	b.Run("get", func(b *testing.B) {
		for range b.N {
			_ = getPointer[SimpleStruct](m)
		}
	})
}

func setPointer[T any](m MP, v VP) {
	m[(*T)(nil)] = v
}

func getPointer[T any](m MP) VP {
	return m[(*T)(nil)]
}
