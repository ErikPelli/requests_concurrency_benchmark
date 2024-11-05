package compute

import (
	"crypto/rand"
	"crypto/sha256"
	"runtime"
	"testing"
)

func shaStaticCoroutines(b *testing.B, n, shaRounds int) {
	channels := make([]chan struct{}, n)
	responses := make([]chan struct{}, n)

	for i := 0; i < n; i++ {
		channels[i] = make(chan struct{})
		responses[i] = make(chan struct{})
		go func(ch, respCh chan struct{}) {
			for range ch {
				buf := make([]byte, _bufSize)
				rand.Read(buf)

				for i := 0; i < shaRounds; i++ {
					sha256.Sum256(buf)
				}
				respCh <- struct{}{}
			}
		}(channels[i], responses[i])
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := range channels {
			channels[i] <- struct{}{}
		}
		for i := range responses {
			<-responses[i]
		}
	}

	// Here I'm not closing the channels, since this is just a short
	// test that will finish immediately after, sorry :)
}

func shaStaticThreads(b *testing.B, n, shaRounds int) {
	channels := make([]chan struct{}, n)
	responses := make([]chan struct{}, n)

	for i := 0; i < n; i++ {
		channels[i] = make(chan struct{})
		responses[i] = make(chan struct{})
		go func(ch, respCh chan struct{}) {
			runtime.LockOSThread()
			for range ch {
				buf := make([]byte, _bufSize)
				rand.Read(buf)

				for i := 0; i < shaRounds; i++ {
					sha256.Sum256(buf)
				}
				respCh <- struct{}{}
			}
		}(channels[i], responses[i])
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for i := range channels {
			channels[i] <- struct{}{}
		}
		for i := range responses {
			<-responses[i]
		}
	}
}
