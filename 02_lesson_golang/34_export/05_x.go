package main

import "fmt"

type MyInterface interface { //Обьявление типа интерфейса
	MethodWithoutParameters()      //Тип поддерживает этот интерфейс если содержит этот метот
	MethodWithParameter(float64)   //а также этот метод с параметром float64
	MethodWithReturnValue() string //и этот метод с вовращаемым значение string
}

type MyType int //Обьявление типа, поддерживающего myinterface

func (m MyType) MethodWithoutParameters() { //Первый обязательный метод.
	fmt.Println("MethodWithoutParameters called")
}

func (m MyType) MethodWithParameter(f float64) { //Второй обязательный метод с параметром float64
	fmt.Println("MethodWithParameter called with", f)
}

func (my MyType) MethodWithReturnValue() string { //Третий обязательный метод с возвращаемым значением string
	return "Hi from MethodWithReturnValue"
}

func (my MyType) MethodNotInterface() { //Тип может поддержиывать интерфейс даже в том случае, если содержит другие методы, не входящие в этот интерфейс
	fmt.Println("MethodNotInterface called")
}

func main() {
	var value MyInterface
	value = MyType(5)
	value.MethodWithoutParameters()
	value.MethodWithParameter(127.3)
	fmt.Println(value.MethodWithReturnValue())
}
