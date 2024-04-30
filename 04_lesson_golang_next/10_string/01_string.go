package main

import "fmt"

func main() {

	name := "Hello world"

	fmt.Println(name)

	// В го строка это набор байтов (в отличии от питона там все что в двойных ковычках строка)
	// Строка - это байтовый слайс со своими особеностями к нижележащему массиву

	//word := "Sample word"
	word := "Тестовая строка" // Сейчас все сломаем

	fmt.Printf("String %s \n", word)
	// Какие значения байт сейчас находятся в слайсе world?
	fmt.Printf("Bytes: ")
	for i := 0; i < len(word); i++ {
		fmt.Printf("%x ", word[i]) // %x - формат предоставления 16-ти ричного байта
	}
	fmt.Println()

	//Каким образом получать доступ к отдельно стоящим символам
	fmt.Printf("Characters: ")
	for i := 0; i < len(word); i++ {
		fmt.Printf("%c ", word[i]) // %c - формат предоставления символа char

	}
	fmt.Println()

	//Строки в го хранятся как наборы utf сиволов и каждый символ можнт занимать больше одного байта
}
