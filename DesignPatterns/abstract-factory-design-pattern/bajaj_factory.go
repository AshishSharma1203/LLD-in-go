package main

import "fmt"

type BajajFactory struct {
}

func NewBajajFactory() *BajajFactory {
	return &BajajFactory{}
}

func (f *BajajFactory) CreateVechile(model VechileModel) (Vechile, error) {
	switch model {
	case ModelBajajPulsar:
		return NewBajajPulsar(), nil

	default:
		return nil, fmt.Errorf("%w:%q (honda factory)", ErrUnkownVechile, model)
	}
}
