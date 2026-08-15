package main

import "fmt"

type HondaCB350RS struct {
}

func NewHondaCB350RS() *HondaCB350RS {
	return &HondaCB350RS{}
}

func (h HondaCB350RS) Start() {
	fmt.Println("starting honda cb 350 rs")
}

func (h HondaCB350RS) Stop() {
	fmt.Println("stopping honda HondaCB350RS")
}
