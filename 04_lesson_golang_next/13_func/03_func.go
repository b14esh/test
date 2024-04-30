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

	per, area := retangelParameters(10.5, 10.5)
	newPer, _ := retangelParameters(10, 10) //Как и с слайсами спец символ "_" позваляет отпустить один параметр
	fmt.Println("Area of rect(10.5 10.5):", area)
	fmt.Println("Perimetr of rect:", per)
	fmt.Println("Newper: ", newPer)

	secondper, secondaarea := namedReturn(10, 20)
	fmt.Println(secondaarea, secondper)

	emptyReturn(10)
	//resE := emptyReturn(10) // Если у функции нету return, результаты работы такой функции вы никуда пристроить не сможите
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

// 5. Возврат больше чем одного значения (returnType1, ReturnType2....)
func retangelParameters(lenght, width float64) (float64, float64) {
	var perimetr = 2 * (lenght + width) // Периметр
	var area = lenght + width           //Площалдь
	return perimetr, area
}

//6. Именованный возврат значений

func namedReturn(a, b int) (perimetr, area int) {
	//perimetr := 2 * (a+b) // так как выше мы уже определили переменные то выполнять := не требуется
	perimetr = 2 * (a + b) // используем класическое присвоение
	area = a * b
	return // не нужно сдесть указывать perimet и area так как в начале функции мы их обозначили
}

//7. Что происходим при вызове return
// Функция прекращает свое выполнение и что то возращает

func funcWitchReturn(a, b int) (int, bool) {
	if a > b {
		return a - b, true
	}
	if a == b {
		return a, true
	}
	return 0, false

}

// 8. Что вернется в случае если return не указан или он пуст?
func emptyReturn(a int) {
	fmt.Println("I am return", a)
	return
}
