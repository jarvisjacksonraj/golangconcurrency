package main

import (
	"fmt"
	"time"
)

// Simulates a worker processing a banking task
func processTask(taskID int) {
	fmt.Printf("Task %d started\n", taskID)
	time.Sleep(time.Duration(taskID) * 100 * time.Millisecond) // Simulate task processing time
	fmt.Printf("Task %d completed\n", taskID)
}

func main() {
	// Without WaitGroup
	for i := 1; i <= 3; i++ {
		go processTask(i)
	}
	time.Sleep(1 * time.Second) // Sleep to ensure tasks complete
	fmt.Println("All tasks completed")
}
