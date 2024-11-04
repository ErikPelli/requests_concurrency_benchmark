package client

import (
	common "github.com/ErikPelli/requests_concurrency_benchmark"
	"runtime"
	"testing"
)

// Create a different client for each request to avoid
// HTTP caches (and HTTP keep-alive), that could affect the timings
// in results.
type cleanClient struct{}

func newCleanClient() *cleanClient {
	return &cleanClient{}
}

func (r *cleanClient) EchoRequest() error {
	return newClient("localhost", common.Port).EchoRequest()
}

func benchmarkEchoRequest(b *testing.B, numClients int, isCoroutine bool) {
	clientFactory := newCleanClient()

	var finalClient echoRequest
	if isCoroutine {
		finalClient = newCoroutineClient(clientFactory, numClients)
	} else {
		// Allow Go to create enough OS threads to handle all the
		// concurrent HTTP requests (+ main goroutine)
		runtime.GOMAXPROCS(numClients + 1)
		finalClient = newThreadClient(clientFactory, numClients)
	}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = finalClient.EchoRequest()
	}
}

func BenchmarkCoroutine_EchoRequest_4(b *testing.B) {
	benchmarkEchoRequest(b, 4, true)
}

func BenchmarkThread_EchoRequest_4(b *testing.B) {
	benchmarkEchoRequest(b, 4, false)
}

func BenchmarkCoroutine_EchoRequest_16(b *testing.B) {
	benchmarkEchoRequest(b, 16, true)
}

func BenchmarkThread_EchoRequest_16(b *testing.B) {
	benchmarkEchoRequest(b, 16, false)
}

func BenchmarkCoroutine_EchoRequest_128(b *testing.B) {
	benchmarkEchoRequest(b, 128, true)
}

func BenchmarkThread_EchoRequest_128(b *testing.B) {
	benchmarkEchoRequest(b, 128, false)
}

func BenchmarkCoroutine_EchoRequest_1024(b *testing.B) {
	benchmarkEchoRequest(b, 1024, true)
}

func BenchmarkThread_EchoRequest_1024(b *testing.B) {
	benchmarkEchoRequest(b, 1024, false)
}

func BenchmarkCoroutine_EchoRequest_2048(b *testing.B) {
	benchmarkEchoRequest(b, 2048, true)
}

func BenchmarkThread_EchoRequest_2048(b *testing.B) {
	benchmarkEchoRequest(b, 2048, false)
}
