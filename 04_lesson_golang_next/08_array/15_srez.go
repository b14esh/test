package main

import (
	"fmt"
)

func main() {

	//Общение со срезом
	// Срез это как массив только без указания начального размера
	slice := []int{10, 20, 30}
	slice[0] = slice[0] * 10
	slice[1] = 200
	slice = append(slice, 200) // Добавление нового значения
	for i, v := range slice {
		fmt.Println(i, v)
	}

	emptySlice := []int{}
	emptySlice = append(emptySlice, 2)
	fmt.Println(emptySlice)
}
