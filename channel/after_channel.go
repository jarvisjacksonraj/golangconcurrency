package main

import (
	"fmt"
	"time"
)

// Simulates sending transaction details for processing
func sendTransactions(ch chan<- string, accountID int, amount float64) {
	transaction := fmt.Sprintf("Transaction for account %d: amount %.2f", accountID, amount)
	time.Sleep(200 * time.Millisecond) // Simulate sending time
	ch <- transaction
	fmt.Printf("Transaction sent for account %d\n", accountID)
}

// Simulates receiving transaction details and processing them
func receiveTransactions(ch <-chan string) {
	for transaction := range ch {
		fmt.Println("Processing", transaction)
		time.Sleep(100 * time.Millisecond) // Simulate processing time
	}
}

func main() {
	// With channels
	ch := make(chan string)
	go sendTransactions(ch, 1, 100.0)
	go sendTransactions(ch, 2, 200.0)
	go sendTransactions(ch, 3, 300.0)
	go receiveTransactions(ch)

	// Sleep to allow all goroutines to finish
	time.Sleep(1 * time.Second)
	close(ch)
}
