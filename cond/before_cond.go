package main

import (
	"fmt"
	"sync"
)

// Simulates a producer generating account transactions
func produceTransactions(transactions *[]int, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Producing transaction %d\n", id)
	*transactions = append(*transactions, id)
}

// Simulates a consumer processing account transactions
func consumeTransactions(transactions *[]int, wg *sync.WaitGroup) {
	defer wg.Done()
	for len(*transactions) > 0 {
		transaction := (*transactions)[0]
		*transactions = (*transactions)[1:]
		fmt.Printf("Consuming transaction %d\n", transaction)
	}
}

func main() {
	var transactions []int
	var wg sync.WaitGroup
	wg.Add(2)
	go produceTransactions(&transactions, 1, &wg)
	go consumeTransactions(&transactions, &wg)
	wg.Wait()
}
