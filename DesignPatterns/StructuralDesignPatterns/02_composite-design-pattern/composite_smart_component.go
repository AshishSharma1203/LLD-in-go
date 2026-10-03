package main

type CompositeSmartComponent struct {
	devices []SmartComponent
}

func (c *CompositeSmartComponent) AddDevice(d SmartComponent) {
	c.devices = append(c.devices, d)
}

func (c *CompositeSmartComponent) TurnOn() {
	for _, component := range c.devices {
		component.TurnOn()
	}
}

func (c *CompositeSmartComponent) TurnOff() {
	for _, component := range c.devices {
		component.TurnOff()
	}
}
