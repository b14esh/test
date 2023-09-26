package main

import "fmt"

func main() {

	notes := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	index := 1
	fmt.Println(index, notes[index]) // Выводит элемент массива с индексом 1.
	index = 3
	fmt.Println(index, notes[index]) // Выводит элемент массива с индексом 3.

	fmt.Println("Hello for next:")
	for i := 0; i <= 2; i++ {
		fmt.Println(i, notes[i])
	}

}
