package main

import "fmt"

type RedLight struct {
	color string
}

func NewRedLight() *RedLight {
	return &RedLight{
		color: "Red",
	}
}

func (g *RedLight) Next(context *TrafficLightContext) {
	fmt.Println("setting next color to Yellow")
	context.SetState(&YellowLight{})
}
func (g *RedLight) Color() string {
	return "RED"
}
