package main

import (
	"fmt"
	"sync"
	"time"
)

type Account struct {
	Balance int
}

var mutex sync.Mutex

func DepositAndWithdraw(account *Account) {
	mutex.Lock()
	defer mutex.Unlock()

	fmt.Println("Balance:", account.Balance)
	if account.Balance < 0 {
		panic(fmt.Sprintf("Balance should not be negative value: %d", account.Balance))
	}
	account.Balance += 1000
	time.Sleep(time.Millisecond)
	account.Balance -= 1000
}

func DepositAndWithdrawWithConcurencyProblem(account *Account) {
	fmt.Println("Balance:", account.Balance)
	if account.Balance < 0 {
		panic(fmt.Sprintf("Balance should not be negative value: %d", account.Balance))
	}
	account.Balance += 1000
	time.Sleep(time.Millisecond)
	account.Balance -= 1000
}

func main() {
	var wg sync.WaitGroup

	account := &Account{Balance: 10}

	fmt.Println("----------Menu----------")
	fmt.Println("Enter 1 to run the code with concurrency problem")
	fmt.Println("Enter 2 to run the code without concurrency problem")

	var input int

	fmt.Scanf("%d", &input)

	if input == 1 {
		wg.Add(10)
		for i := 0; i < 10; i++ {
			go func() {
				for {
					DepositAndWithdrawWithConcurencyProblem(account)
				}
			}()
		}
		wg.Wait()
	} else {
		wg.Add(10)
		for i := 0; i < 10; i++ {
			go func() {
				for {
					DepositAndWithdraw(account)
				}
			}()
		}
		wg.Wait()
	}

}
