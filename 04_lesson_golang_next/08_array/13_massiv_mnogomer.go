package main

import "fmt"

func main() {

	//Много мерные моссивы со значениями
	//Может содержать массивы одного типа (только int или только string или только float64 или etc.)
	words := [2][2]string{
		{"BOB", "Alice"},
		{"Vasia", "Petr"},
	}

	fmt.Println("Multidimensional array:", words)

}
