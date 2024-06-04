package main

import (
	"fmt"
	"sync"
)

// Simulates a producer generating account transactions
func produceTransactions(cond *sync.Cond, transactions *[]int, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	cond.L.Lock()
	fmt.Printf("Producing transaction %d\n", id)
	*transactions = append(*transactions, id)
	cond.Signal()
	cond.L.Unlock()
}

// Simulates a consumer processing account transactions
func consumeTransactions(cond *sync.Cond, transactions *[]int, wg *sync.WaitGroup) {
	defer wg.Done()
	cond.L.Lock()
	for len(*transactions) == 0 {
		cond.Wait()
	}
	transaction := (*transactions)[0]
	*transactions = (*transactions)[1:]
	fmt.Printf("Consuming transaction %d\n", transaction)
	cond.L.Unlock()
}

func main() {
	var transactions []int
	var mu sync.Mutex
	cond := sync.NewCond(&mu)
	var wg sync.WaitGroup
	wg.Add(2)
	go produceTransactions(cond, &transactions, 1, &wg)
	go consumeTransactions(cond, &transactions, &wg)
	wg.Wait()
}
