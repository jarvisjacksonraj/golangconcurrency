package main

import (
	"fmt"
	"time"
)

var accountBalance float64 // Shared resource

// Simulates updating the account balance
func updateBalance(amount float64) {
	fmt.Printf("Updating balance by %.2f\n", amount)
	accountBalance += amount
	time.Sleep(100 * time.Millisecond) // Simulate processing time
	fmt.Printf("New balance: %.2f\n", accountBalance)
}

func main() {
	// Without mutex
	for i := 1; i <= 5; i++ {
		go updateBalance(float64(i) * 100.0)
	}
	time.Sleep(1 * time.Second) // Sleep to ensure all updates complete
	fmt.Printf("Final balance: %.2f\n", accountBalance)
}
