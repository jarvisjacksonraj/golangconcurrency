package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Account struct {
	ID      int
	Balance int
	Mutex   sync.Mutex
}

var accounts = make(map[int]*Account)
var accountIDCounter int
var mutex sync.Mutex
var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go handleCommands()
	defer wg.Wait()

	scanner := bufio.NewScanner(os.Stdin)
	displayMenu() // Show options at the start
	for scanner.Scan() {
		command := scanner.Text()
		if command == "exit" {
			break
		}
		executeCommand(command)
		displayMenu() // Show options after each command
	}
}

func handleCommands() {
	defer wg.Done()
	// Implement command handling logic here
}

func createAccount(initialBalance int) int {
	mutex.Lock()
	defer mutex.Unlock()
	accountIDCounter++
	accounts[accountIDCounter] = &Account{ID: accountIDCounter, Balance: initialBalance}
	return accountIDCounter
}

func deposit(accountID, amount int) {
	if account, exists := accounts[accountID]; exists {
		account.Mutex.Lock()
		account.Balance += amount
		account.Mutex.Unlock()
		fmt.Printf("Deposited %d to account %d\n", amount, accountID)
	} else {
		fmt.Println("Account not found.")
	}
}

func withdraw(accountID, amount int) {
	if account, exists := accounts[accountID]; exists {
		account.Mutex.Lock()
		if account.Balance >= amount {
			account.Balance -= amount
			fmt.Printf("Withdrew %d from account %d\n", amount, accountID)
		} else {
			fmt.Println("Insufficient funds.")
		}
		account.Mutex.Unlock()
	} else {
		fmt.Println("Account not found.")
	}
}

func transfer(fromID, toID, amount int) {
	fromAccount, fromExists := accounts[fromID]
	toAccount, toExists := accounts[toID]

	if fromExists && toExists {
		fromAccount.Mutex.Lock()
		defer fromAccount.Mutex.Unlock()
		toAccount.Mutex.Lock()
		defer toAccount.Mutex.Unlock()

		if fromAccount.Balance >= amount {
			fromAccount.Balance -= amount
			toAccount.Balance += amount
			fmt.Printf("Transferred %d from account %d to account %d\n", amount, fromID, toID)
		} else {
			fmt.Println("Insufficient funds for transfer.")
		}
	} else {
		fmt.Println("One or both accounts not found.")
	}
}

func displayAllAccounts() {
	fmt.Println("\nAccount Details:")
	for id, account := range accounts {
		account.Mutex.Lock()
		fmt.Printf("Account ID: %d, Balance: %d\n", id, account.Balance)
		account.Mutex.Unlock()
	}
}

func displayMenu() {
	fmt.Println("\nAvailable Commands:")
	fmt.Println("  create <initialBalance>  - Create a new account with a specified initial balance.")
	fmt.Println("  deposit <accountID> <amount> - Deposit an amount into a specific account.")
	fmt.Println("  withdraw <accountID> <amount> - Withdraw an amount from a specific account.")
	fmt.Println("  transfer <fromAccountID> <toAccountID> <amount> - Transfer an amount from one account to another.")
	fmt.Println("  display - Display all account balances.")
	fmt.Println("  exit - Exit the program.")
	fmt.Println("Please enter a command:")
}

func executeCommand(command string) {
	parts := strings.Fields(command)
	if len(parts) == 0 {
		return
	}

	switch parts[0] {
	case "create":
		if len(parts) < 2 {
			fmt.Println("Usage: create <initialBalance>")
			return
		}
		initialBalance, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("Invalid amount.")
			return
		}
		id := createAccount(initialBalance)
		fmt.Printf("Account %d created with balance %d.\n", id, initialBalance)
	case "deposit", "withdraw", "transfer":
		if len(parts) < 3 {
			fmt.Printf("Usage: %s <accountID> <amount>\n", parts[0])
			return
		}
		accountID, err1 := strconv.Atoi(parts[1])
		amount, err2 := strconv.Atoi(parts[2])
		if err1 != nil || err2 != nil {
			fmt.Println("Invalid command parameters.")
			return
		}
		if parts[0] == "deposit" {
			deposit(accountID, amount)
		} else if parts[0] == "withdraw" {
			withdraw(accountID, amount)
		} else if parts[0] == "transfer" {
			if len(parts) < 4 {
				fmt.Println("Usage: transfer <fromAccountID> <toAccountID> <amount>")
				return
			}
			toID, err := strconv.Atoi(parts[3])
			if err != nil {
				fmt.Println("Invalid to account ID.")
				return
			}
			transfer(accountID, toID, amount)
		}
	case "display":
		displayAllAccounts()
	default:
		fmt.Println("Unknown command.")
	}
}
