package main

import "fmt"

func main() {
	// coffee:=NewEspresso()
	var coffee Coffee = NewEspresso()
	coffee = NewMilkDecorator(coffee)
	coffee = NewSugarDecorator(coffee)

	fmt.Println("order is :", coffee.GetDescription())

}
