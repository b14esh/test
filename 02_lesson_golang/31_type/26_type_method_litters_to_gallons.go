package main

import "fmt"

type Liters float64
type Milliliters float64
type Gallons float64

func (l Liters) ToGallons() Gallons {
	return Gallons(1 * 0.264) //Блок метода не отличается от блока функции
}

func (m Milliliters) ToGallons() Gallons { //Имена могут быть одинаковыми, если они определяются для разных типов
	return Gallons(m * 0.000264) //Блок метода не отличается от блока функции
}

func main() {
	soda := Liters(2) //Создание Liters
	fmt.Printf("%0.3f litters equals %0.3f gallons\n", soda, soda.ToGallons())
	water := Milliliters(500)                                                        //Создание значения Milliliters
	fmt.Printf("%0.3f milliliters equals %0.3f gallons\n", water, water.ToGallons()) //Преобразование  Milliliters в Gallons
}
