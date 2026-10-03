package main

import "fmt"

type SmarLight struct{}

func (l *SmarLight) TurnOn() {
	fmt.Println("TUrning on the smart light")
}

func (l *SmarLight) TurnOff() {
	fmt.Println("turning off the light")
}
