package compute

import (
	"crypto/rand"
	"crypto/sha256"
	"runtime"
	"sync"
)

const _bufSize = 256

func shaCoroutines(n, shaRounds int) {
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			buf := make([]byte, _bufSize)
			rand.Read(buf)

			for i := 0; i < shaRounds; i++ {
				sha256.Sum256(buf)
			}
		}()
	}
	wg.Wait()
}

func shaThreads(n, shaRounds int) {
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			runtime.LockOSThread()
			defer wg.Done()

			buf := make([]byte, _bufSize)
			rand.Read(buf)

			for i := 0; i < shaRounds; i++ {
				sha256.Sum256(buf)
			}
		}()
	}
	wg.Wait()
}
