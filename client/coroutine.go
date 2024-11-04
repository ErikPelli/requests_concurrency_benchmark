package client

import (
	"log"
	"sync"
)

type coroutine struct {
	cli      echoRequest
	requests int
}

func newCoroutineClient(cli echoRequest, requests int) *coroutine {
	return &coroutine{
		cli:      cli,
		requests: requests,
	}
}

func (c *coroutine) EchoRequest() error {
	var wg sync.WaitGroup
	for i := 0; i < c.requests; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := c.cli.EchoRequest(); err != nil {
				log.Fatal("unexpected error", err)
			}
		}()
	}
	wg.Wait()

	return nil
}
