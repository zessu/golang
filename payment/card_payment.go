package main

import (
	"errors"
	"fmt"
	"regexp"
	"time"
)

// CreateCreditAccount creates a new credit account
func CreateCreditAccount(ownerName string, cardNumber string, expirationMonth int, expirationYear int, securityCode int) *CreditCard {
	return &CreditCard{
		ownerName:       ownerName,
		cardNumber:      cardNumber,
		expirationMonth: expirationMonth,
		expirationYear:  expirationYear,
		securityCode:    securityCode,
	}
}

// OwnerName returns name of card owner, this is a getter
func (c *CreditCard) OwnerName() string {
	return c.ownerName
}

// SetOwnerName sets a cards owner property to the name of the owner, this is a setter
func (c *CreditCard) SetOwnerName(value string) error {
	if len(value) == 0 {
		return errors.New("Invalid Owner Name provided")
	}
	c.ownerName = value
	return nil
}

// CardNumber returns the card number for any given card
func (c *CreditCard) CardNumber() string {
	return c.cardNumber
}

var cardNumberPattern = regexp.MustCompile("\\d{4}-\\d{4}-\\d{4}-\\d{4}")

// SetCardNumber sets a cards number
func (c *CreditCard) SetCardNumber(value string) error {
	if !cardNumberPattern.Match([]byte(value)) {
		return errors.New("Invalid Card number")
	}
	c.cardNumber = value
	return nil
}

// SetExpirationDate sets expiration date for a credit card
func (c *CreditCard) SetExpirationDate(year int, month int) error {
	now := time.Now()
	if year < now.Year() || (year == now.Year() && time.Month(month) < now.Month()) {
		return errors.New("Expiration date not valid")
	}
	c.expirationMonth, c.expirationYear = month, year
	return nil
}

// SecurityCode returns the security code for this card
func (c *CreditCard) SecurityCode() int {
	return c.securityCode
}

// AvailableCredit returns amount left in card
func (c *CreditCard) AvailableCredit() float32 {
	return c.availableCredit
}

// ProcessPayment processes a credit card payment
func (c *CreditCard) ProcessPayment() bool {
	fmt.Println("processing card payment")
	return true
}
