package main 

type CarBuilder struct{
	engine            string
	brand             string
	model             string
	year              int
	color             string
	sunroof           bool
	navigationSystem  bool
}
// car builder with default values 

func NewCarBuilder() *CarBuilder {
	return &CarBuilder{
		year:            2020,
		color:           "White",
		sunroof:        false,
		navigationSystem: false,
	}
}

func (b *CarBuilder) SetEngine(engine string) *CarBuilder {
	b.engine = engine
	return b
}

func (b *CarBuilder) SetBrand(brand string) *CarBuilder {
	b.brand = brand
	return b
}

func (b *CarBuilder) SetModel(model string) *CarBuilder {
	b.model = model
	return b
}

func (b *CarBuilder) SetYear(year int) *CarBuilder {
	b.year = year
	return b
}

func (b *CarBuilder) SetColor(color string) *CarBuilder {
	b.color = color
	return b
}

func (b *CarBuilder) SetSunroof(sunroof bool) *CarBuilder {
	b.sunroof = sunroof
	return b
}

func (b *CarBuilder) SetNavigationSystem(navigationSystem bool) *CarBuilder {
	b.navigationSystem = navigationSystem
	return b
}

func (b *CarBuilder) Build() *Car {
	return newCar(b)
}
