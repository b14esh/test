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

func main() {

	//fmt.Println("hello world")

	//4.Вызов простейшей функции
	res := add(10, 20)
	fmt.Println("Result of add(10, 20):", res)
	fmt.Println("Result of mult(1,2,3,4):", mult(1, 2, 3, 4))

}

//3.Простейшая функция (определить как до ммомента ее вызова в функции main,
// так и в любом месте пакета, главное что бы лна была определена в принципе где-нибудь)

func add(a int, b int) int {
	result := a + b
	return result
}

// 4. Функция с однатипными параметрами
func mult(a, b, c, d int) int {
	result := a * b * c * d
	return result
}
