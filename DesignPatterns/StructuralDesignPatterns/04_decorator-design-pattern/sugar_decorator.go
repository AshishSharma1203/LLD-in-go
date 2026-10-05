package main

type SugarDecorator struct {
	coffee Coffee
}

func NewSugarDecorator(coffee Coffee) *SugarDecorator {
	return &SugarDecorator{coffee: coffee}
}

func (s *SugarDecorator) GetDescription() string {
	return s.coffee.GetDescription() + " ,  Sugar"
}

func (s *SugarDecorator) GetCost() float64 {
	return s.coffee.GetCost() + 0.25
}
