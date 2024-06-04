package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Simulates incrementing a transaction counter atomically
func incrementCounter(counter *int32, wg *sync.WaitGroup) {
	defer wg.Done()
	atomic.AddInt32(counter, 1)
	fmt.Printf("Counter value: %d\n", atomic.LoadInt32(counter))
}

func main() {
	var counter int32
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go incrementCounter(&counter, &wg)
	}
	wg.Wait()
	fmt.Println("Final counter value:", atomic.LoadInt32(counter))
}
