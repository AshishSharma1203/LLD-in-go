package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	vechile1, err := NewVechile(VechileTypeCar)
	if err != nil {
		log.Fatal(err)
	}
	vechile1.Start()
	vechile1.Stop()

	vechile2, err := NewVechile(VechileTypeBike)
	if err != nil {
		log.Fatal(err)
	}
	vechile2.Start()
	vechile2.Stop()

	_, err = NewVechile(VechileType("ship"))

	if err != nil {
		if errors.Is(err, ErrUnkownVechile) {
			fmt.Println("expected error:", err)
		} else {
			log.Fatal(err)
		}
	}

}
