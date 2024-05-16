package main

import "fmt"

func main() {

	fmt.Println("Иницилизация в блоке услоного операторая")
	//for ( int i = 0; i < 0; i++){}

	// Блок присваивания для if только вот так вот ":="
	if num := 10; num%2 == 0 {
		fmt.Println("EVEN", num)
	} else {
		fmt.Println("ODD", num)
	}
}
