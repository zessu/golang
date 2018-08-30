package main

import (
	"fmt"
)

// CreateCashAccount creates a cash account
func CreateCashAccount() *Cash {
	return &Cash{}
}

// ProcessPayment processes a cash payment
func (c *Cash) ProcessPayment() bool {
	fmt.Println("processing cash payment")
	return true
}
