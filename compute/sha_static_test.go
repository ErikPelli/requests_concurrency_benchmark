package compute

import (
	"runtime"
	"testing"
)

func BenchmarkShaStaticCoroutines(b *testing.B) {
	const n = 6
	shaStaticCoroutines(b, n, 10000)
}

func BenchmarkShaStaticThreads(b *testing.B) {
	const n = 6
	runtime.GOMAXPROCS(n + 1)

	shaStaticThreads(b, n, 10000)
}
