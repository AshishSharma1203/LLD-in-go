package main

import "fmt"

type Car struct {
	engine           string
	brand            string
	model            string
	year             int
	color            string
	sunroof          bool
	navigationSystem bool
}

func newCar(b *CarBuilder) *Car {
	return &Car{
		engine:           b.engine,
		brand:            b.brand,
		model:            b.model,
		year:             b.year,
		color:            b.color,
		sunroof:          b.sunroof,
		navigationSystem: b.navigationSystem,
	}
}

func (c *Car) Engine() string {
	return c.engine
}

func (c *Car) Brand() string {
	return c.brand
}

func (c *Car) Model() string {
	return c.model
}

func (c *Car) Year() int {
	return c.year
}

func (c *Car) Color() string {
	return c.color
}

func (c *Car) Sunroof() bool {
	return c.sunroof
}

func (c *Car) NavigationSystem() bool {
	return c.navigationSystem
}

func (c *Car) String() string {
	return fmt.Sprintf("Car [engine=%s, brand=%s, model=%s, year=%d, color=%s, sunroof=%t, navigationSystem=%t]",
		c.engine, c.brand, c.model, c.year, c.color, c.sunroof, c.navigationSystem)
}
