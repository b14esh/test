package main

import "fmt"

type Gallons float64
type Liters float64
type Milliliters float64

func (g Gallons) ToLiters() Liters { //Определяет метод ToLiters для nbgf Пфддщты
	return Liters(g * 3.785)
}

func (g Gallons) ToMilliliters() Milliliters { // Определяет метод ToMilliliters для типа Gallons
	return Milliliters(g * 3785.41)
}

func main() {
	milk := Gallons(2)
	fmt.Printf("%0.3f gallons equals %0.3f liters\n", milk, milk.ToLiters())
	fmt.Printf("%0.3f gallons equals %0.3f milliliters\n", milk, milk.ToMilliliters())
}
