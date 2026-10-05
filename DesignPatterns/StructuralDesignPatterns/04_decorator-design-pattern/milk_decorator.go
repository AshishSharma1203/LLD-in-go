package main

type MilkDecorator struct {
	coffee Coffee
}

func NewMilkDecorator(coffee Coffee) *MilkDecorator {
	return &MilkDecorator{coffee: coffee}
}

func (m *MilkDecorator) GetDescription() string {
	return m.coffee.GetDescription() + ", Milk"
}

func (m *MilkDecorator) GetCost() float64 {
	return m.coffee.GetCost() + 0.50
}
