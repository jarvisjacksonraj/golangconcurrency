package main

import (
	"fmt"
	"time"
)

// Simulates sending transaction details for processing
func sendTransactions(accountID int, amount float64) {
	fmt.Printf("Sending transaction for account %d: amount %.2f\n", accountID, amount)
	time.Sleep(200 * time.Millisecond) // Simulate sending time
	fmt.Printf("Transaction sent for account %d\n", accountID)
}

// Simulates processing transactions sequentially
func processTransactionsSequentially() {
	sendTransactions(1, 100.0)
	sendTransactions(2, 200.0)
	sendTransactions(3, 300.0)
}

func main() {
	// Without channels
	fmt.Println("Starting transaction processing")
	processTransactionsSequentially()
	fmt.Println("Completed transaction processing")
}
