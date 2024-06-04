package main

import (
	"fmt"
	"time"
)

// Simulates processing a banking transaction
func processTransaction(accountID int, amount float64) {
	fmt.Printf("Processing transaction for account %d: amount %.2f\n", accountID, amount)
	time.Sleep(100 * time.Millisecond) // Simulate time taken to process
	fmt.Printf("Transaction completed for account %d\n", accountID)
}

func main() {
	// Without goroutine
	fmt.Println("Before processing transactions")
	processTransaction(1, 100.0)
	processTransaction(2, 200.0)
	processTransaction(3, 300.0)
	fmt.Println("After processing transactions")
}
