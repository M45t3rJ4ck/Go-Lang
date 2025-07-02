package main

import (
	"fmt"
	"sync"
	"time"
)

// Concurrent task in Go is called Go Routines.
func printTo15() {
	for i := 1; i <= 15; i++ {
		fmt.Println("Function 1:", i)
	}
}

func printTo10() {
	for i := 1; i <= 10; i++ {
		fmt.Println("Function 2:", i)
	}
}

// Communicate via channels
func nums1(channel chan int) {
	channel <- 1
	channel <- 2
	channel <- 3
}

func nums2(channel chan int) {
	channel <- 4
	channel <- 5
	channel <- 6
}

// Bank Account Operation Simulation
type Account struct {
	balance float64
	// Mutex = Mutual exclusion
	lock sync.Mutex
}

func (a *Account) GetBalance() float64 {
	// Lock operation so no other function uses this.
	a.lock.Lock()
	// Unlock so other function can use this again.
	defer a.lock.Unlock()

	return a.balance
}

func (a *Account) Withdraw(v float64) {
	a.lock.Lock()
	defer a.lock.Unlock()

	if v > a.balance {
		fmt.Println("Not enugh money in account.")
	} else {
		fmt.Printf("%.2f Withdrawn : %.2f Balance\n", v, a.balance)
		a.balance -= v
	}
}

func main() {
	// Concurrently execute multiple functions - Functions can't have return values.
	go printTo15()
	go printTo10()
	time.Sleep(2 * time.Second)

	// Communicate via channels
	channel1 := make(chan int)
	channel2 := make(chan int)

	go nums1(channel1)
	go nums2(channel2)

	fmt.Println(<-channel1)
	fmt.Println(<-channel2)
	fmt.Println(<-channel1)

	fmt.Println(<-channel2)
	fmt.Println(<-channel1)
	fmt.Println(<-channel2)

	// Bank Account Operation Simulation
	var acct Account

	acct.balance = 100

	fmt.Println("Balance: ", acct.GetBalance())

	for i := 0; i < 12; i++ {
		go acct.Withdraw(10)
	}
	time.Sleep(2 * time.Second)
}
