package main

import "fmt"

type YamahaFactory struct {
}

func NewYamahaFactory() *YamahaFactory {
	return &YamahaFactory{}
}

func (f *YamahaFactory) CreateVechile(model VechileModel) (Vechile, error) {
	switch model {
	case ModelYamahaXSR:
		return NewYamahaXSR155(), nil

	default:
		return nil, fmt.Errorf("%w:%q (honda factory)", ErrUnkownVechile, model)
	}
}
