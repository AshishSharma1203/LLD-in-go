package main

import "fmt"

type Car struct {
}

func NewCar() (car *Car) {
	return &Car{}
}

func (c *Car) Start() {
	fmt.Println("Car is starting")
}

func (c *Car) Stop() {
	fmt.Println("Car is stopping")
}
