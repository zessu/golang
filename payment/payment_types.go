package main

// CreditCard represents a credit card
type CreditCard struct {
	ownerName       string
	cardNumber      string
	expirationMonth int
	expirationYear  int
	securityCode    int
	availableCredit float32
}

// Cash represents a payment type
type Cash struct {
	balance float32
	amount  float32
}

// PaymentOption applies for all cash methods
type PaymentOption interface {
	ProcessPayment() bool
	checkBalance() float32
}
