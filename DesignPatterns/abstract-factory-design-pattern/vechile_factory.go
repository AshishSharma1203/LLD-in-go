package main

import (
	"errors"
	"fmt"
)

type Brand string

const (
	BrandHonda  Brand = "Honda"
	BrandYamaha Brand = "Yamaha"
	BrandBajaj  Brand = "Bajaj"
)

type VechileModel string

const (
	ModelHondaCB50RS   VechileModel = "honda-cb-350rs"
	ModelHondaRebel300 VechileModel = "honda-rebel-3500"
	ModelBajajPulsar   VechileModel = "bajaj-pulsar"
)

const (
	ModelYamahaXSR VechileModel = "yamaha-xsr-155"
)

var ErrUnkownBrand = errors.New("unknown brand")

var ErrUnkownVechile = errors.New("unknown vechile")

type VechileFactory interface {
	CreateVechile(model VechileModel) (Vechile, error)
}

func NewVechileFactory(brand Brand) (VechileFactory, error) {
	switch brand {

	case BrandHonda:
		return NewHondaFactory(), nil
	case BrandYamaha:
		return NewYamahaFactory(), nil
	case BrandBajaj:
		return NewBajajFactory(), nil

	default:
		return nil, fmt.Errorf("%w %q:", ErrUnkownBrand, brand)
	}
}
