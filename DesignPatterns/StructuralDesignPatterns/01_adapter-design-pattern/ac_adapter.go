package main

import "fmt"

type AirConditionerAdapter struct {
	airConditioner *AirConditioner
}

func NewAirConditionerAdapter(ac *AirConditioner) *AirConditionerAdapter {
	return &AirConditionerAdapter{airConditioner: ac}
}

func (ac *AirConditionerAdapter) TurnOn() {
	fmt.Println("turnign on the ac via adapter")
	ac.airConditioner.connectToBluetooth()
	ac.airConditioner.startCooling()
}

func (ac *AirConditionerAdapter) TurnOff() {
	fmt.Println("turnign off the ac via adapter")

	ac.airConditioner.stopCooling()
	ac.airConditioner.disconnectBluetooth()
}
