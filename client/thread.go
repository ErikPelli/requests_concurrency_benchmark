package client

import (
	"log"
	"runtime"
	"sync"
)

type thread struct {
	cli      echoRequest
	requests int
}

func newThreadClient(cli echoRequest, requests int) *thread {
	return &thread{
		cli:      cli,
		requests: requests,
	}
}

func (t *thread) EchoRequest() error {
	var wg sync.WaitGroup
	for i := 0; i < t.requests; i++ {
		wg.Add(1)
		go func() {
			// LockOSThread associates the goroutine to an OS thread
			// (This is the only difference)
			// Documentation:
			// - The calling goroutine will always execute in that thread,
			//    and no other goroutine will execute in it.
			// - If the calling goroutine exits without unlocking the thread,
			//    the thread will be terminated.
			runtime.LockOSThread()

			defer wg.Done()
			if err := t.cli.EchoRequest(); err != nil {
				log.Fatal("unexpected error", err)
			}
		}()
	}
	wg.Wait()

	return nil
}
