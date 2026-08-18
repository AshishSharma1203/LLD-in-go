package main

import "fmt"

type PaymentProcessor struct {
	PaymentMethod Payment
}

func NewPaymentProcessor(PaymentMethod Payment) *PaymentProcessor {
	fmt.Printf("creating payment processor of payment ")
	return &PaymentProcessor{
		PaymentMethod: PaymentMethod,
	}
}

func (p *PaymentProcessor) SetStrategy(PaymentMethod Payment) {
	p.PaymentMethod = PaymentMethod
}

func (p *PaymentProcessor) ProcessPayment(amount float32) {
	p.PaymentMethod.Pay(amount)
}
