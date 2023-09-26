package main

import "fmt"

func main() {
	notes := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	primes := [5]int{2, 3, 5, 7, 11}

	fmt.Println(notes[4], primes[1])

	text := [3]string { // Все это один массив.
		"This is a series of long strings",
		"which would be awkward to place",
		"together on a single line", // Запятая в конце обязательна.
		}
	fmt.Println(text[2])

	fmt.Println(notes)
	fmt.Println(primes)

	fmt.Printf("%#v\n", notes)
	fmt.Printf("%#v\n", primes)
}
