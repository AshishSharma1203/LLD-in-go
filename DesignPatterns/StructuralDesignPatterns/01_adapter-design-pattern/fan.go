package main

import "fmt"

type Fan struct {
}

func (b *Fan) turnOn() {
	fmt.Println("turning on the fan")
}

func (b *Fan) turnOff() {
	fmt.Println("turning off the fan")
}

func (b *Fan) speedUp() {
	fmt.Println("speeding up the fan ")
}
