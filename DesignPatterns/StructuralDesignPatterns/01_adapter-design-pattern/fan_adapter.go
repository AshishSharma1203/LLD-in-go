package main

import "fmt"

type FanAdapter struct {
	fan *Fan
}

func NewFanAdapter(fan *Fan) *FanAdapter {
	return &FanAdapter{fan: fan}
}

func (fAdp *FanAdapter) TurnOn() {
	fmt.Println("turnign on the fan via adapter")

	fAdp.fan.turnOn()
	fAdp.fan.speedUp()
}

func (fAdp *FanAdapter) TurnOff() {
	fmt.Println("turnign off the fan ia adapter")

	fAdp.fan.turnOff()
}
