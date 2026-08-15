package main

import "fmt"

type HondaFactory struct {
}

func NewHondaFactory() *HondaFactory {
	return &HondaFactory{}
}

// func (f *HondaFactory) CreateVechile(model VechileModel) (Vechile, error) {
// 	switch model {
// 	case ModelHondaCB50RS:
// 		return NewHondaCB350RS(), nil

// 	case ModelHondaRebel300:
// 		return NewHondaRebel300(), nil
// 	default:
// 		return nil, fmt.Errorf("%w:%q (honda factory)", ErrUnkownVechile, model)
// 	}
// }

func (f *HondaFactory) CreateVechile(model VechileModel) (Vechile, error) {
	switch model {
	case ModelHondaCB50RS:
		return NewHondaCB350RS(), nil

	case ModelHondaRebel300:
		return NewHondaRebel300(), nil
	default:
		return nil, fmt.Errorf("%w:%q (honda factory)", ErrUnkownVechile, model)
	}
}

