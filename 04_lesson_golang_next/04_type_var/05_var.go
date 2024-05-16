package main

import "fmt"

func main() {

	//var объявляет 1 или более переменных.
	var a string = "initial"
	fmt.Println(a)

	//Вы можете объявить несколько переменных за раз.
	var b, c int = 1, 2
	fmt.Println(b, c)

	//Go определит тип инициализированных переменных автоматически.

	var d = true
	fmt.Println(d)

	//Не инициализированные переменные имеют нулевое значение. Например нулевое значение для int это 0.
	var e int
	fmt.Println(e)

	//Синтаксис := это сокращение для объявления и инициализации переменной, например в этом примере var f string = "short".
	f := "short"
	fmt.Println(f)

	//обьявление пачки переменных
	var (
		g int = 77
		h int = 1
		z int = g + h
	)

	fmt.Printf("Hello g=%d, h=%d, z=g+h=%d\n", g, h, z)

}
