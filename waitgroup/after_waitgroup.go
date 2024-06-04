package main

import (
	"fmt"
	"sync"
	"time"
)

// Simulates a worker processing a banking task
func processTask(taskID int, wg *sync.WaitGroup) {
	defer wg.Done() // Signal that this goroutine is done
	fmt.Printf("Task %d started\n", taskID)
	time.Sleep(time.Duration(taskID) * 100 * time.Millisecond) // Simulate task processing time
	fmt.Printf("Task %d completed\n", taskID)
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1) // Add a goroutine to the WaitGroup
		go processTask(i, &wg)
	}
	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("All tasks completed")
}
