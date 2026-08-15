package main

import (
	"errors"
	"fmt"
)

type VechileType string

const (
	VechileTypeCar   VechileType = "car"
	VechileTypeBike  VechileType = "bike"
	VechileTypeTruck VechileType = "truck"
)

var ErrUnkownVechile = errors.New("unknown vechile type")

func NewVechile(vechileType VechileType) (Vechile, error) {
	switch vechileType {
	case VechileTypeCar:
		return NewCar(), nil

	case VechileTypeBike:
		return NewBike(), nil

	case VechileTypeTruck:
		return NewTruck(), nil

	default:
		return nil, fmt.Errorf("%w:%q", ErrUnkownVechile, vechileType)
	}
}
