/*
Программа должна объявлять целочисленную переменную myInt и переменную
myIntPointer для хранения целочисленного указателя. Затем она должна присваивать
значение myInt и указатель на myInt как значение myIntPointer. Наконец, программа
должна выводить значение по указателю myIntPointer.
*/

package main

import "fmt"

func main() {
	var myInt  int = 100
	var myIntPointer int = 200
	fmt.Println(myInt)
	fmt.Println(myIntPointer)
	x := &myInt
	*x = 150
	fmt.Println(myInt)
	y := &myIntPointer
	*y = 800
	fmt.Println(myIntPointer)
}

//Результат = 42
