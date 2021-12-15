package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	start := time.Now().UnixNano()
	for i := 0; i < 2048; i++ {
		wg.Add(1)
		// index := i
		go func() {
			fib(33)
			// fmt.Printf("[Goroutine:%d] fib(32)=%d \n", index, fib(32))
			wg.Done()
		}()
	}
	wg.Wait()

	d := time.Now().UnixNano() - start
	ms := float64(d) / float64(time.Millisecond)

	fmt.Printf("\nAll goroutines done: %v ms \n", ms)
}

func fib(n uint64) uint64 {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
