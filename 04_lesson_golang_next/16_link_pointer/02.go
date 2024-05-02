package main

import "fmt"

func changeParam(val *int) {
	//8.1 Определение функции, принимающией параметр как указатель
	*val += 100
}

func returnPointer() *int {
	//9.1 Иницилизация функции возвращающий указатель
	var numeric int = 371
	return &numeric // В момент возврата GO перемещает данныую переменную в кучу
}

func main() {
	//8. Передача указателей в функцию
	//Коласальный прирост производительности из за того что передается не значение
	//а передается лишь адрес в памяти за которым хранится значение
	sample := 1
	//samplePointer := &sample
	fmt.Println("Origin value of sample:", sample)
	//changeParam(samplePointer)
	changeParam(&sample)
	fmt.Println("After changing samle is:", sample)

	//9. Возврат поинтера из функции (В C++ результат работы такого механизма - неопределен)
	ptr1 := returnPointer()
	ptr2 := returnPointer()
	fmt.Printf("Ptr1: %Tand address %v and value %v\n", ptr1, ptr1, *ptr1)
	fmt.Printf("Ptr2: %Tand address %v and value %v\n", ptr2, ptr2, *ptr2)
}
