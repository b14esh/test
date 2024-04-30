package main

import "fmt"

func main() {

	//Создание массива через иницилизацию переменных
	arrWithValue := [...]int{10, 20, 30, 40}
	// Узнать длину массива используем len(name_arr)
	fmt.Println("arrWithValue:", arrWithValue, "Lenght:", len(arrWithValue))
}
