package main

import (
	"fmt"
	"reflect"
)

func main() {
	var myInt int
	fmt.Println(reflect.TypeOf(&myInt)) //Получает указатель на myInt и выводит тип указателя.
	var myFloat float64
	fmt.Println(reflect.TypeOf(&myFloat)) //Получает указатель на myFloat и выводит тип указателя.
	var myBool bool
	fmt.Println(reflect.TypeOf(&myBool)) //Получает указатель на myBool и выводит тип указателя.*
}
