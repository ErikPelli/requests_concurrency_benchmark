package compute

import (
	"runtime"
	"testing"
)

func BenchmarkShaCoroutines(b *testing.B) {
	const n = 32
	for i := 0; i < b.N; i++ {
		shaCoroutines(n, 100)
	}
}

func BenchmarkShaThreads(b *testing.B) {
	const n = 32
	runtime.GOMAXPROCS(n + 1)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		shaThreads(n, 100)
	}
}
