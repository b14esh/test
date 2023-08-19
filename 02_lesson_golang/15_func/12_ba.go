package main

import "fmt"

func main() {

	var myInt int
	var myIntPointer *int //Объявление переменной, содержит указатель int
	myIntPointer = &myInt //Указатель присваивается переменной
	fmt.Println(myIntPointer)
	var myFloat float64
	var myFloatPointer *float64 //Объявление переменной для хранения указателя на float64
	myFloatPointer = &myFloat   //Указатель присваивается переменной
	fmt.Println(myFloatPointer)

}
