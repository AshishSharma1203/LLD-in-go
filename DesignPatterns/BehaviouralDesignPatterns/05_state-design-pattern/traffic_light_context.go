package main

type TrafficLightContext struct {
	currentState TrafficLightState
}

func NewTrafficLightContext() *TrafficLightContext {
	return &TrafficLightContext{
		currentState: &RedLight{},
	}
}

func (context *TrafficLightContext) SetState(state TrafficLightState) {
	context.currentState = state
}

func (context *TrafficLightContext) Next() {
	context.currentState.Next(context)
}
