package main

import (
	"fmt"
)

func main() {

	fmt.Println("Иницилизация в блоке услоного операторая")
	//for ( int i = 0; i < 0; i++){}

	// Блок присваивания для if только вот так вот ":="
	fmt.Println("Введите число")
	var xxx int
	fmt.Scan(&xxx)

	if num := xxx; num%2 == 0 {
		fmt.Println("EVEN", num)
	} else {
		fmt.Println("ODD", num)
	}
}
