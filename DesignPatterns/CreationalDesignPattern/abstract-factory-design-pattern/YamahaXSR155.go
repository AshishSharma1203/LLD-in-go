package main

import "fmt"

type YamahaXSR155 struct {
}

func NewYamahaXSR155() *YamahaXSR155 {
	return &YamahaXSR155{}
}

func (h YamahaXSR155) Start() {
	fmt.Println("starting honda cb 350 rs")
}

func (h YamahaXSR155) Stop() {
	fmt.Println("stopping honda YamahaXSR155")
}
