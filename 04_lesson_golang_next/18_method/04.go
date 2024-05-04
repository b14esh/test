//Point_VS_VAL

package main

import "fmt"

type Rectangle struct {
	length, width int
}

//1. Реализуем функцию и метод подсчета периметра прямоугольника
// Важно: все параметры передаем как копии

//4. При вызове этого метода неважно, будет ли он вызыватся у экземпляра или у его ссылки
func (r Rectangle) Perimeter() int {
	return 2 * (r.length + r.width)
}

//5. данную функцию можно вызвать только у копии прямоугольника (но не у его ссылки)
func Perimeter(r Rectangle) int {
	return 2 * (r.length + r.width)
}

//6. Допустим будет метод который меняет значения состояние структуры на новые но этот метот ValueResive
func (r Rectangle) SetLength(newLength int) {
	r.length = newLength
}

func main() {
	//2. Создаем экземпляр прямоугольника
	rectangleASVAule := Rectangle{10, 10}
	fmt.Println("Call function rentangleASValue", Perimeter(rectangleASVAule))
	fmt.Println("Call function rentangleASValue", rectangleASVAule.Perimeter())

	//3. Создадим указатель на прямоугольник
	rectangasPointer := &rectangleASVAule
	fmt.Println("Call method for rentangleASValue", rectangasPointer.Perimeter())
	//Perimeter(rectangleASVAule) //error // Элюстрация к пункту 5
	//.\04.go:31:22: cannot use &rectangleASVAule (value of type *Rectangle) as Rectangle value in assignment

	//7. Вызываем метод SetLength у экземпляра
	fmt.Println("Befor call method SetLength", rectangleASVAule)
	rectangleASVAule.SetLength(9999)
	fmt.Println("After call method SetLength", rectangleASVAule)

	//8. Вызываем метод SetLength у ссылки на rectangleASValuse
	rectangasPointer.SetLength(99999999)
	fmt.Println("After call method &rectangle", *rectangasPointer)

}
