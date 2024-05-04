package main

import "fmt"

type Rectangle struct {
	length, width int
}

func (r *Rectangle) Area() int {
	return r.length * r.width
}

func Area(r *Rectangle) int {
	return r.length * r.width
}

func (r *Rectangle) SetWidth(newWidth int) {
	r.width = newWidth
}

func main() {

	rectangleAsValue := Rectangle{10, 10}
	rectangleAsPointer := &rectangleAsValue
	fmt.Println("call are function for &renctangle", Area(rectangleAsPointer))
	fmt.Println("call are function for &renctangle", rectangleAsPointer.Area())

	//1. Вызываем метод у value - исходное значение
	fmt.Println("call are function for renctangle", rectangleAsPointer.Area())
	//2. Вызов функции с параметром value - исходное значение
	//Area(rectangleAsValue) //err
	//.\05.go:27:7: cannot use rectangleAsValue (variable of type Rectangle) as *Rectangle value in argument to Area

	//3. Рампечатаем исходный прямоугольник
	fmt.Println("Before changing wifth:", rectangleAsValue)

	//4 Вызываю метот у serWidth у &rentangle
	rectangleAsPointer.SetWidth(9999)
	fmt.Println("After changing via method SetWidth fo &rentangle:", rectangleAsValue)

	//5. Вызов метода у rectangle
	rectangleAsValue.SetWidth(888)
	fmt.Println("After changing via method SetWidth fo rentangle:", rectangleAsValue)
}
