package main

import "fmt"

func main() {
	myInt := 4 //Значение
	myIntPointer := &myInt //Указатель
	fmt.Println(myIntPointer) //Выводится сам указатель.
	fmt.Println(*myIntPointer) //Выводится значение, на которое ссылается указатель.

	myFloat := 98.6 //Значение
	myFloatPointer := &myFloat //Указатель
	fmt.Println(myFloatPointer) //Выводит сам указатель
	fmt.Println(*myFloatPointer) //Выводится значение, на которое ссылается указатель.

	myBool := true //Значение
	myBoolPointer := &myBool     ////Указатель
	fmt.Println(myBoolPointer)   //Выводит сам указатель
	fmt.Println(*myBoolPointer) //Выводится значение, на которое ссылается указатель.
}
