package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	HondaFactory, err := NewVechileFactory(BrandHonda)
	if err != nil {
		log.Fatal(err)
	}
	cb350, err := HondaFactory.CreateVechile(ModelHondaCB50RS)
	if err != nil {
		log.Fatal(err)
	}
	cb350.Start()
	cb350.Stop()
	rebel300, err := HondaFactory.CreateVechile(ModelHondaRebel300)
	if err != nil {
		log.Fatal(err)
	}
	rebel300.Start()
	rebel300.Stop()

	// unkown model for this brand
	if _, err := HondaFactory.CreateVechile(ModelBajajPulsar); err != nil {
		if errors.Is(err, ErrUnkownVechile) {
			fmt.Println("expected err :", err)
		} else {
			log.Fatal(err)
		}
	}

	if _, err := NewVechileFactory(Brand("tesla")); err != nil {
		if errors.Is(err, ErrUnkownBrand) {
			fmt.Println("expected err: ", err)
		} else {
			log.Fatal(err)
		}
	}
}
