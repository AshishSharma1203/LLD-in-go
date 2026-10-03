package main

import "fmt"

type SmartFan struct{}

func (fan *SmartFan) TurnOn() {
	fmt.Println("trning on the fan")
}

func (fan *SmartFan) TurnOff() {
	fmt.Println("turning off the fan")
}
