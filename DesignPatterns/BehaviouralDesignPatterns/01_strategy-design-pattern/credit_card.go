package main

import "fmt"

type CreditCard struct {
	name string
}

func NewCreditCard(name string) *CreditCard {
	return &CreditCard{
		name: name,
	}
}

func (c CreditCard) Pay(amount float32) {
	fmt.Printf("processing amount  %f using method %s", amount, c.name)
}
