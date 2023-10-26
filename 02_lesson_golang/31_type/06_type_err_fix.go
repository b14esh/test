package main

import "fmt"

type Gallons float64
type Liters float64


func ToGallons(l Liters) Gallons { //Значение типа Gallons чуть больше 1/4 от значения типа Liters.
	return Gallons(l * 0.264)
}
func ToLiters(g Gallons) Liters { //Значение типа Liters почти вчетверо больше значения типа Gallons.
	return Liters(g * 3.785)
}
func main() {
	carFuel := Gallons(1.2)
	busFuel := Liters(4.5)
	carFuel += ToGallons(Liters(40.0)) //Liters преобразуется в Gallons перед сложением.
	busFuel += ToLiters(Gallons(30.0)) //Gallons преобра-зуется в Liters перед сложением.
	fmt.Printf("Car fuel: %0.1f gallons\n", carFuel)
	fmt.Printf("Bus fuel: %0.1f liters\n", busFuel)
}
