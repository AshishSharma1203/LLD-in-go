package main

type Espresso struct {
}

func NewEspresso() *Espresso {
	return &Espresso{}
}

func (e *Espresso) GetDescription() string {
	return "Espresso"
}

func (e *Espresso) GetCost() float64 {
	return 3.00
}
