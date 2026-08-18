package main

import "fmt"

type Truck struct {
}

func NewTruck() (truck *Truck) {
	return &Truck{}
}

func (c *Truck) Start() {
	fmt.Println("Truck is starting")
}

func (c *Truck) Stop() {
	fmt.Println("Truck is stopping")
}
