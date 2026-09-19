// Shared memory ko safe rakhna: Mutex aur atomic.
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// SafeCounter: map/int ko multiple goroutines se touch karna ho to lock zaroori hai
type SafeCounter struct {
	mu sync.Mutex
	n  map[string]int
}

func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	defer c.mu.Unlock() // defer se unlock bhoolne ka chance nahi
	c.n[key]++
}

func main() {
	c := SafeCounter{n: map[string]int{}}
	var atomicN atomic.Int64 // simple counter ke liye atomic mutex se tez hai
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc("hits")
			atomicN.Add(1)
		}()
	}
	wg.Wait()
	fmt.Println(c.n["hits"], atomicN.Load()) // 1000 1000

	// tip: `go run -race ./mutex` se data race pakad sakte ho
}
