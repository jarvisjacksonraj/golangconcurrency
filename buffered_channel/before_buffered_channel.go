package main

import (
	"fmt"
	"time"
)

// Simulates sending account balance updates
func updateBalance(accountID int, amount float64, ch chan<- float64) {
	fmt.Printf("Updating balance for account %d: amount %.2f\n", accountID, amount)
	time.Sleep(150 * time.Millisecond) // Simulate update time
	ch <- amount
	fmt.Printf("Balance updated for account %d\n", accountID)
}

// Simulates receiving and processing balance updates
func processBalanceUpdates(ch <-chan float64) {
	for amount := range ch {
		fmt.Printf("Processing balance update: amount %.2f\n", amount)
		time.Sleep(100 * time.Millisecond) // Simulate processing time
	}
}

func main() {
	// Without buffered channels
	ch := make(chan float64)
	go updateBalance(1, 100.0, ch)
	go updateBalance(2, 200.0, ch)
	go updateBalance(3, 300.0, ch)
	go processBalanceUpdates(ch)

	// Sleep to allow all goroutines to finish
	time.Sleep(1 * time.Second)
	close(ch)
}
