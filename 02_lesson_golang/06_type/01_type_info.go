package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Написал целое число 42, go пределили  тип:", reflect.TypeOf(42))
	fmt.Println("Написал дробное число 3.1234, go определил тип:", reflect.TypeOf(3.1234))
	fmt.Println("Написал слово true и не занес в двойные ковычки, go определил тип:", reflect.TypeOf(true))
	fmt.Println("Написал фразу в двойных ковычках \"Hello world \", go определил тип:", reflect.TypeOf("Hello world"))
}
