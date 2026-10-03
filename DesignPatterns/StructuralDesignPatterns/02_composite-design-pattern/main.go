package main

func main(){
	light1:=&SmarLight{}
	light2:=&SmarLight{}

	fan1:=&SmartFan{}
	fan2:=&SmartFan{}

	room1:=&CompositeSmartComponent{}
	room2:=&CompositeSmartComponent{}

	room1.AddDevice(fan1)
	room1.AddDevice(light1)

	room2.AddDevice(fan2)
	room2.AddDevice(light2)

	room1.TurnOn()
	room2.TurnOn()

	floor1:=&CompositeSmartComponent{}
	floor2:=&CompositeSmartComponent{}
	floor1.AddDevice(room1)
	floor2.AddDevice(room2)

	house:=&CompositeSmartComponent{}
	house.AddDevice(floor1)
	house.AddDevice(floor2)

	house.TurnOn()

}