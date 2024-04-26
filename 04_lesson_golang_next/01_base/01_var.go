package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("Hello world")
	// - комментарий

	/*
		Коментарий внутри
	*/

	//////////////////////////////
	//////////////////////////////

	// декларирование переменных
	var age int
	fmt.Println("My age is:", age)
	age = 32
	fmt.Println("Age after assignment: ", age)

	//Декларирование и иницилизация пользовательских значений
	var height int = 186
	fmt.Println("My height is: ", height)

	//В чем "полустрогость" типизации? Можно опускать тип переменной
	var uid = 123456
	fmt.Println("My uid", uid)

	// Как узнать тип переменной?
	fmt.Printf("%T\n", uid)

	//Декларирование и иницилизация переменных одного типа

	//var firstVar, secondVar = 20,30
	var firstVar, secondVar int = 20, 30
	fmt.Printf("\n firstVar:%d  Secondvar:%d\n", firstVar, secondVar)

	//Декларирование блока переменных

	var (
		personName string = "Bob"
		persongAge int    = 42
		personUid  int
	)

	fmt.Printf("\nName: %s \nAge: %d  \nUid: %d\n\n", personName, persongAge, personUid)

	//Немного странного // избегайте и создавайте в виде блока
	var a, b = 30, "Baris"
	fmt.Println(a, b)

	//Немного хорошего // Повторное деклорирование приведет к ошибке
	//var a = 200
	//Присвоение работает без проблем
	a = 400

	//Коротка деклорация // Не рекомендуется использовать
	//var count int = 10
	// Вот так вот
	count := 10
	fmt.Println(count)
	count++
	fmt.Println(count)
	count = count + 1
	fmt.Println(count)

	//Множественные присваивания через :=
	//aArg, bArg := 10, "Boba" - хз у меня это не работает
	//# command-line-arguments
	//.\01_var.go:70:20: undefined: BOB
	// Возможно в новых версия go это больше не работает
	// создавайте блок переменных и избегите ошибку

	var (
		aArg int    = 10
		bArg string = "Boba"
	)
	fmt.Println("\n", aArg, bArg)

	x1, x2 := 10, 20
	fmt.Println(x1, x2)
	// x1,x2 := 20, 40 не получится так как переменные уже задекларированы
	//  но есть исключение, если к сущетвующей добавть новую переменную ио все получится
	x2, x3 := 44, 66
	fmt.Println(x1, x2, x3)

	//Пример
	width, lenght := 20.5, 30.2
	fmt.Printf("Min dimensional of rectangle is : %.3f", math.Min(width, lenght))
}
