package main

import (
	"fmt"
	"sync"
)

var initialized bool
var accountNumber int

// Simulates initializing the account number
func initializeAccountNumber() {
	fmt.Println("Initializing account number...")
	accountNumber = 123456
	initialized = true
}

func getAccountNumber() int {
	if !initialized {
		initializeAccountNumber()
	}
	return accountNumber
}

func main() {
	// Without once
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Printf("Account number: %d\n", getAccountNumber())
		}()
	}
	wg.Wait()
}
