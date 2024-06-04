package main

import (
	"fmt"
	"sync"
)

// Simulates storing and accessing account balances
func updateAccountBalance(accountBalances map[int]float64, accountID int, amount float64, wg *sync.WaitGroup) {
	defer wg.Done()
	accountBalances[accountID] += amount
	fmt.Printf("Updated balance for account %d: %.2f\n", accountID, accountBalances[accountID])
}

func main() {
	accountBalances := make(map[int]float64)
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go updateAccountBalance(accountBalances, i, float64(i)*100.0, &wg)
	}
	wg.Wait()
	fmt.Println("Final account balances:", accountBalances)
}
