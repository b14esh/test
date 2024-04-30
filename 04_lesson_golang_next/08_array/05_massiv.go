package main

import "fmt"

func main() {

	//Создание массива через и не указывать элементы // Что будет?
	arrWithValue := [...]int{}
	// Узнать длину массива используем len(name_arr)
	fmt.Println("arrWithValue:", arrWithValue, "Lenght:", len(arrWithValue))
	// Длина будет 0
	// Увеличить длину массива после иницилизации не льзя !
	arrWithValue[0] = 10000
	fmt.Println("arrWithValue:", arrWithValue, "Lenght:", len(arrWithValue))
	// ошибка CreateFile .\05_massive.go: The system cannot find the file specified.
}
