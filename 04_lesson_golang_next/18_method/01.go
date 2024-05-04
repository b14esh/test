package main

import (
	"fmt"
	"math"
)

//3. Вчем преимущество методов над функциями
// Во-пурвых наличие методов улучает "консистентность" кода, т.к напрямую влияет на его понимание
// Во-вторых: в Go запрещается создовать функции с одинаковыми названиями, в то время как методы для разных струтур,
// содинаковыми названиями - разрешены

type Circle struct {
	radius float64
}

type Rectangle struct {
	lenght, width int
}

// 4. Создадим функцию метод Perimetr для структуры Circle
// выкручиваемся с помощью функции
func PerimeterCircle(c Circle) float64 {
	return c.radius * 2 * math.Pi

}

func Perimeter(c Circle) float64 {
	return c.radius * 2 * math.Pi
}

func (c Circle) Perimeter() float64 {
	return c.radius * 2 * math.Pi

}

// 5. Создадим функцию и метод Perimeter для структуры Rentangle
func (r Rectangle) Perimeter() int {
	return 2 * (r.lenght + r.width)
}

//func Perimeter(r Rectangle) int {
//	return 2 * (r.lenght + r.width)
//	}
//\01.go:36:6: Perimeter redeclared in this block
//.\01.go:22:6: other declaration of Perimeter

// 6. В го рарешено создавать методы с одинаковыми именами в пределах одной структуры.
// Главное чтобы получатель метода в разных структурах (где реализован метод со схожим методом) оличался.
func PerimeterRentagle(r Rectangle) int {
	return 2 * (r.lenght + r.width)
}

func main() {
	c := Circle{10.5}
	fmt.Println("call function", Perimeter(c))
	fmt.Println("call function", PerimeterCircle(c))

	fmt.Println("call method", c.Perimeter())

	r := Rectangle{10, 20}
	fmt.Println("Call funcrion for rentangle", PerimeterRentagle(r))
	fmt.Println("Call method for rentagle", r.Perimeter())

}
