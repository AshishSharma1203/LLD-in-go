package main

import "fmt"

type Bike struct {
}
func NewBike() (bike *Bike) {
	return &Bike{}
}
func (c *Bike) Start() {
	fmt.Println("Bike is starting")
}

func (c *Bike) Stop() {
	fmt.Println("Bike is stopping")
}
