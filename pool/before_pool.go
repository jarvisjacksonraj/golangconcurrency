package main

import (
	"fmt"
	"sync"
)

// Simulates a database connection
type Connection struct {
	ID int
}

// Simulates creating a new database connection
func createConnection(id int) *Connection {
	fmt.Printf("Creating connection %d\n", id)
	return &Connection{ID: id}
}

func main() {
	// Without pool
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			conn := createConnection(id)
			fmt.Printf("Using connection %d\n", conn.ID)
		}(i)
	}
	wg.Wait()
}
