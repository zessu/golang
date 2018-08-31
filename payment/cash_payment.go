package main

import (
	"fmt"
)

// CreateCashAccount creates a cash account
func CreateCashAccount(amt float32) *Cash {
	return &Cash{
		amount: amt,
	}
}

// ProcessPayment processes a cash payment
func (c *Cash) ProcessPayment() bool {
	fmt.Println("processing cash payment")
	return true
}

// Balance returns amount left in cash
func (c *Cash) checkBalance() float32 {
	return c.amount
}
