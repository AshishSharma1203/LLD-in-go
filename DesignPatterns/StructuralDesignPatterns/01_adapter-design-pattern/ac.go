package main

import "fmt"

type AirConditioner struct {
}

func (ac *AirConditioner) stopCooling() {
	fmt.Println("stopping the cooling")
}

func (ac *AirConditioner) connectToBluetooth() {
	fmt.Println("connecting to the bluetooth")
}

func (ac *AirConditioner) disconnectBluetooth() {
	fmt.Println("disconect from bluetooth")
}

func (ac *AirConditioner) startCooling() {
	fmt.Println("starting the cooling ")
}
