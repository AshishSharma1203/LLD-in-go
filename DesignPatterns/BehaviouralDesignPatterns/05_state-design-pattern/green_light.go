package main

import "fmt"

type GreenLight struct {
	color string
}

func NewGreenLight() *GreenLight {
	return &GreenLight{
		color: "Green",
	}
}

func (g *GreenLight) Next(context *TrafficLightContext) {
	fmt.Println("setting next color to Red")
	context.SetState(&RedLight{})
}
func (g *GreenLight) Color() string {
	return "GREEN"
}
