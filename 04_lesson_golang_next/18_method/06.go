package main

import "fmt"

//method for standarts

//1. Методы для стандартных типов
//В Go встроено куча примитивов (int, int32, string, bool ...)
//Что если очень хочется дописать к стандартному типу какой-то метод?

//2. Наивная попытка. Это не выполнимо.
// Компелятор запрещает добавление методов к сущетвуещим базовым типам
// cannot define new methods on non-local type int
// https://pkg.go.dev/golang.org/x/tools/internal/typesinternal#InvalidRecv

/*
func (a *int) isEven() bool {
	if *a%2 == 0 {
		return true
	}
	return false
}
*/

//3. Но мнре очень хочется что делать ?
// создайте новый тип - ваш int и делайте  ним что хотите!
// Для создания нового типа используют конструкцию

type MySuperDuperInt int

func (msdi *MySuperDuperInt) IsEven() bool {
	if *msdi%2 == 0 {
		return true
	}
	return false
}

func main() {
	num1 := MySuperDuperInt(202)
	num2 := MySuperDuperInt(201)
	fmt.Println(num1.IsEven())
	fmt.Println(num2.IsEven())
	//4. Внутренние преобразования
	// var num3 MySuperDuperInt = MySuperDuperInt(10)
	// var num4 int = int(num3)

}
