package main

import (
	"fmt"
	"sync"
)

// Simulates incrementing a transaction counter
func incrementCounter(counter *int32, wg *sync.WaitGroup) {
	defer wg.Done()
	*counter++
	fmt.Printf("Counter value: %d\n", *counter)
}

func main() {
	var counter int32
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go incrementCounter(&counter, &wg)
	}
	wg.Wait()
	fmt.Println("Final counter value:", counter)
}
