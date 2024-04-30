package main

import "fmt"

func main() {

	//Конкантенация строк (сложение строк)
	//word1, word2 := "Вася", "Петя"
	//word3 := word1 + " " + word2
	//fmt.Println(word3)

	//Строитель строк (java -> stringbuilder)

	firstname := "Alex"
	secondname := "Jonson"
	fullname := fmt.Sprintf("%s #### %s", firstname, secondname) // fmt.Sprintf() - некая функция форматирования иногда бывает полезна и позволяет форматировать что угодно
	fmt.Println("Fullname:", fullname)

	//СТРОКИ НЕ ИЗМЕНЯЕМЫЕ  // ИМУТАБЕЛЬНЫЕ

	// А слайсы изменяемые

	fullnameSlice := []rune(fullname)
	fullnameSlice[0] = 'F'
	fullname = string(fullnameSlice)
	fmt.Println("String mutation", fullname)

	//Сравнение рун
	//if 'Q' == 'Q' {
	if 'Q' == 'F' {
		fmt.Println("rune equal")
	} else {
		fmt.Println("rune  not equal")
	}

	// Где живут полезные методы работы со строкамаи?
	// import "strings"

}
