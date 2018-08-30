package main

import (
	"fmt"
)

func processPayment(option PaymentOption) {
	option.ProcessPayment()
}

func main() {
	card := CreateCreditAccount(
		"john dhoe",
		"1111-2222-3333-4444",
		6,
		2018,
		123,
	)
	cash := CreateCashAccount()
	fmt.Println(card.OwnerName())
	fmt.Println(card.CardNumber())
	error := card.SetCardNumber("1111-1111-1111-1111")
	if error == nil {
		var option PaymentOption
		fmt.Println(card.CardNumber())
		option = card // use abstracted away functionality
		option.ProcessPayment()
		option = cash // change option to instance of cash which also implements the interface
		option.ProcessPayment()
		// or we can do this using a function instead
		processPayment(card)
		processPayment(cash)
	} else {
		fmt.Println("There was an error changing that card number")
	}
}
