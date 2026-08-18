package main

import "fmt"

func main() {

	builder1 := NewCarBuilder()
	car1 := builder1.
		SetBrand("Honda").
		SetColor("Red").
		SetEngine("V8").
		SetNavigationSystem(true).
		SetSunroof(true).
		Build()

	fmt.Println(car1)

	// always create a new builder for each car as using prev ones have the same values as previous car
	builder2 := NewCarBuilder()
	car2 := builder2.
		SetBrand("Toyota").
		SetColor("Blue").
		SetEngine("V6").
		SetNavigationSystem(false).
		SetSunroof(false).
		Build()
	
	fmt.Println(car2)

}
