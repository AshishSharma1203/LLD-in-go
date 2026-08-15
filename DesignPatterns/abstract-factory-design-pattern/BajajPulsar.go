package main

import "fmt"

type BajajPulsar struct {
}

func NewBajajPulsar() *BajajPulsar {
	return &BajajPulsar{}
}

func (h BajajPulsar) Start() {
	fmt.Println("starting honda cb 350 rs")
}

func (h BajajPulsar) Stop() {
	fmt.Println("stopping honda BajajPulsar")
}
