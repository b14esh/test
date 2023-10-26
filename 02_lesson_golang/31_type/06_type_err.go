package main

import "fmt"

type Liters float64
type Gallons float64

func main() {
	carFuel := Gallons(1.2)
	busFuel := Liters(2.5)
	carFuel += Liters(8.0)   //Значение Liters нель-зя сложить со значением Gallons!
	busFuel += Gallons(30.0) //Значение Gallons нель-зя сложить со значением Liters!
}
