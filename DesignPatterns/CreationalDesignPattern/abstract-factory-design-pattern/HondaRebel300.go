package main

import "fmt"

type HondaRebel300 struct {
}

func NewHondaRebel300() *HondaRebel300 {
	return &HondaRebel300{}
}

func (h HondaRebel300) Start() {
	fmt.Println("starting honda HondaRebel300")
}

func (h HondaRebel300) Stop() {
	fmt.Println("stopping honda HondaRebel300")
}
