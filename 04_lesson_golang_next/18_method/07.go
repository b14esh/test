//constructor

package main

import "fmt"

//1. Создатим структуру Rectangle
type Rectangle struct {
	length, width int
}

//2. Создадим конструктор для Rectangle
//Для имен конструкторов в Go договорились использовать функциб с следующим названием:
// * New() если данный конструктор на файл один (в файле присутствует описание одной структурой)
// * New<StructName() - если в файле присутсвует еще какие-то структуры

//3. В го принято возращать из конструктора не сам экземпляр, а ссылку на него

func NewRectangle(newLength, newWidth int) *Rectangle {
	return &Rectangle{newLength, newWidth}
}

//4. Добавим два метода

// Периметр
func (r *Rectangle) Perimeter() int {
	return 2 * (r.length + r.width)
}

// Плошадь
func (r *Rectangle) Area() int {
	return r.length * r.width
}

type Circle struct {
	radius float64
}

func NewCircle(newRadius float64) *Circle {
	return &Circle{newRadius}

}

func main() {
	r := NewRectangle(10, 20)
	fmt.Println(r)
	fmt.Printf("Type as %T and value %v\n", r, r)
	fmt.Println("Perimeter:", r.Perimeter(), "Area", r.Area())
	fmt.Println()
}
