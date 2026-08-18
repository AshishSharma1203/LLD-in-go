package main

import "fmt"

type UPIPayment struct {
	name string
}

func NewUPIPayment(name string) *UPIPayment {
	return &UPIPayment{
		name: name,
	}
}

func (u UPIPayment) Pay(amount float32) {
	fmt.Printf("doing payment via upi %s of amount %v", u.name, amount)
}
