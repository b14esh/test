/*
Программа должна объявлять целочисленную переменную myInt и переменную
myIntPointer для хранения целочисленного указателя. Затем она должна присваивать
значение myInt и указатель на myInt как значение myIntPointer. Наконец, программа
должна выводить значение по указателю myIntPointer.
*/

package main

import "fmt"

func main() {
	var myInt int = 42
	myIntPointer := &myInt
	fmt.Println(*myIntPointer)
}

//Результат = 42
