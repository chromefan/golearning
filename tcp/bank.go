// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 261.
//!+

// Package bank provides a concurrency-safe bank with one account.
package main

import (
	"fmt"
	"sync"
)

var (
	mu      sync.Mutex // guards balance
	balance int
)

func Deposit(amount int) {
	mu.Lock()
	balance = balance + amount
	mu.Unlock()
}
func Expend(amount int) {
	balance = balance - amount
}
func Balance() int {
	mu.Lock()
	b := balance
	mu.Unlock()
	return b
}

//!-

func main() {
	// Deposit [1..1000] concurrently.
	var n sync.WaitGroup
	for i := 1; i <= 10; i++ {
		n.Add(1)
		go func(amount int) {
			Deposit(amount)
			Expend(amount)
			fmt.Printf("amount %d\n", amount)
			n.Done()
		}(i)
		go func(amount int) {
			Expend(amount)
			fmt.Printf("Expend %d\n", amount)
			n.Done()
		}(i%3)
	}
	n.Wait()
	b := Balance()
	fmt.Printf(" Balance %d", b)
}
