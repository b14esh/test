package main

import (
	"fmt"
)

// 1. Явная функция - локально определенный блок кода, имеющий имя(ЯВНОЕ ОПРЕДЕЛЕНИЕ)
// Функцию необходимо определить
// Функцию необходимо вызвать
// 2. Сигнатура функции
//
//	func functionName(params type) TypeReturnValue {
//		//body
//	}
func functionWO() {

}

//3.Простейшая функция

func add(a int, b int) int {
	result := a + b
	return result
}

func main() {

	//fmt.Println("hello world")

	//4.Вызов простейшей функции
	res := add(10, 20)
	fmt.Println("Result of add(10, 20):", res)

}
