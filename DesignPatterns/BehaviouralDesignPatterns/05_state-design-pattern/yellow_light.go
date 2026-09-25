package main

import "fmt"

type YellowLight struct {
	color string
}

func NewYellowLight() *YellowLight {
	return &YellowLight{
		color: "Yellow",
	}
}

func (g *YellowLight) Next(context *TrafficLightContext) {
	fmt.Println("setting next color to Green")
	context.SetState(&GreenLight{})
}
func (g *YellowLight) Color() string {
	return "YELLOW"
}
