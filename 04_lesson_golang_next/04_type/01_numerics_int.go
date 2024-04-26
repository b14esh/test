package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// Nemerics. Integers
	// int8, int16, int32, int64, int
	// uint8, uint16, uint32, uint64, uint
	// 8bytes = 64 bit
	var a int = 32
	b := 92
	fmt.Println("Value of a:", a, "Value of b:", b, "Sum of a+b:", a+b)
	c := a + b
	// Вывод типа через %T форматирование
	fmt.Printf("Type a is %T\n", a)
	fmt.Printf("Type b is %T\n", b)
	fmt.Printf("Type c is %T\n", c)
	// Узнаем сколько байт занимает переменная типа int
	fmt.Printf("a is type %T and size type %d \n", a, unsafe.Sizeof(a))
	fmt.Printf("b is type %T and size type %d \n", b, unsafe.Sizeof(b))
	fmt.Printf("c is type %T and size type %d \n", c, unsafe.Sizeof(c))
	// При использовании короткого обьявлния - тип для целого числа - int платформа зависимый (то есть для 64bit системы это будет int64 (это равно 8Byte если проверять unsafe.Sizeof())

	//Эксперемент
	var first32 int32 = 12
	var second64 int64 = 13
	//.\01_numerics_int.go:30:14: invalid operation: first32 + second64 (mismatched types int32 and int64)
	//fmt.Println(first32 + second64)
	// С помощью int64 или int32 мы можем преобразовать число в нужную битность и потом выпольнить произведение
	fmt.Println(first32 + int32(second64))
	fmt.Printf("\n%d \n", int64(first32)+second64)
	// Используйте явное привидение типов при необходимости если уверены что не произойдет коллизии

	//Эксперемент
	var az64 int64 = 16123414
	var azint int = 156234
	//fmt.Println(az64 + azint)
	//err .\01_numerics_int.go:40:14: invalid operation: az64 + azint (mismatched types int64 and int)
	// Если проводятся арифметические операции, то оббязательно нужно производить механизм привидения. Т.к. int != int64
	fmt.Println()
	fmt.Println(az64 + int64(azint))

	// + - * / %
	// % остаток от деления

	// Аналогичным образом устроены uint8, uint16, uint32, uint64, uint
	// uint* хранит исключительно положительные числа!
}
