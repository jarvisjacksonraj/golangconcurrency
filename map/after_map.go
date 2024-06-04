package main

import (
	"fmt"
	"sync"
)

// Simulates storing and accessing account balances with synchronization
func updateAccountBalance(mu *sync.Mutex, accountBalances map[int]float64, accountID int, amount float64, wg *sync.WaitGroup) {
	defer wg.Done()
	mu.Lock()
	accountBalances[accountID] += amount
	fmt.Printf("Updated balance for account %d: %.2f\n", accountID, accountBalances[accountID])
	mu.Unlock()
}

func main() {
	accountBalances := make(map[int]float64)
	var mu sync.Mutex
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go updateAccountBalance(&mu, accountBalances, i, float64(i)*100.0, &wg)
	}
	wg.Wait()
	mu.Lock()
	fmt.Println("Final account balances:", accountBalances)
	mu.Unlock()
}
