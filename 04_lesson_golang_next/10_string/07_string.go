package main

import "fmt"

func main() {

	//Сравнение строк
	// Нормально начало работать == и != с go 1.6

	word1, word2 := "Вася", "Петя"

	if word1 == word2 { // равно ли word1  word2
		fmt.Println("Equal")
	} else {
		fmt.Println("Not equal")
	}

	if word1 != word2 { //
		fmt.Println("Equal")
	} else {
		fmt.Println("Not equal")
	}

	// Со знаком > или < могут возникнуть проблемы, так как там идет сравнение по байтам, и одни символ может весть 4 байта, а другой например 1 байт....
}
