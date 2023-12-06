// package mypackage - могли бы создать отдельный пакет
package main

import "fmt"

type MyType struct { //MyTypeОбьявляем как тип сруктуры
	EmbeddedType // EmbeddedType встраивается в MyType
}

type EmbeddedType string //Обьявление встраиваемого типа (неважно, является ли он структурой)

func (e EmbeddedType) ExportedMethod() { //Этот метод будет повышен до MyType
	fmt.Println("Hi from ExportedMethod on EmbeddedType")
}

func (e EmbeddedType) unexportedMethod() { //Этот метод повышаться не будет
}

/*
И могли бы замутить его вызов из другого пакета
package main
import "mypackage"
*/

func main() {
	//value := mypackage.MyType{} вызывали бы мы из пакета mypackage
	value := MyType{}
	value.ExportedMethod() //Вызывается метод, повышенный с уровня EmbeddedType
	//value.unexportedMethod() при попытки вызвать это, будет ошибка
}
