package benchmarks

import "testing"

type VI = string
type KI = any
type MI = map[KI]VI

func BenchmarkIntKeys(b *testing.B) {
	m := make(MI)
	b.Run("set", func(b *testing.B) {
		for i := range b.N {
			setInt(m, i, "hello world")
		}
	})

	b.Run("get", func(b *testing.B) {
		for i := range b.N {
			_ = getInt(m, i)
		}
	})
}

func setInt(m MI, k KI, v VI) {
	m[k] = v
}

func getInt(m MI, k KI) VI {
	return m[k]
}
