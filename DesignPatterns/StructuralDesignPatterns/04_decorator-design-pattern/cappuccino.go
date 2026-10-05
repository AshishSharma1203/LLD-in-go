package main

type Cappucino struct {
}

func NewCappucino() *Cappucino {
	return &Cappucino{}
}

func (c *Cappucino) GetDescription() string {
	return "Cappucino"
}

func (c *Cappucino) GetCost() float64 {
	return 3.00
}
