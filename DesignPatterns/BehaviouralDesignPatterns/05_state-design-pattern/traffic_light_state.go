package main

type TrafficLightState interface {
	Next(context *TrafficLightContext)
	Color() string
}
