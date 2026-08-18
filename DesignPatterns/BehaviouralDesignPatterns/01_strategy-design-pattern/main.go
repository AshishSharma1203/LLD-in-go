package main

func main() {
	paymentMethod1 := NewCreditCard("ICICI Amex card")
	paymentMethod2 := NewUPIPayment("CredPAY")

	paymentProcessor := NewPaymentProcessor(paymentMethod1)
	paymentProcessor.ProcessPayment(3000.12)
	paymentProcessor.SetStrategy(paymentMethod2)
	paymentProcessor.ProcessPayment(131351.123)
}
