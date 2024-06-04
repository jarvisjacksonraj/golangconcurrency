package main

import (
	"fmt"
	"sync"
)

var once sync.Once
var accountNumber int

// Simulates initializing the account number
func initializeAccountNumber() {
	fmt.Println("Initializing account number...")
	accountNumber = 123456
}

func getAccountNumber() int {
	once.Do(initializeAccountNumber)
	return accountNumber
}

func main() {
	// With once
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
