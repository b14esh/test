package main

import "fmt"

type MyType string //Определяется новый тип

func (m MyType) sayHi() {
	fmt.Println("Hi")
}

func main() {
	value := MyType("a MyType value") //Создается значение MyType.
	value.sayHi() //Вызывается sayHi для этого значения
	anotherValue := MyType("another value") //Создается другое значение MyType
	anotherValue.sayHi() //Вызывается sayHi для нового значения.
}
