package main

import "fmt"

func main() {
	fmt.Println("Runing code for adapter design pattern")
	ac := &AirConditioner{}
	ac1 := NewAirConditionerAdapter(ac)
	fan := &Fan{}
	fan1 := NewFanAdapter(fan)

	ac1.TurnOn()
	fan1.TurnOn()
	fan1.TurnOff()
	fan1.TurnOff()
}
