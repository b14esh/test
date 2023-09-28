package main

import "fmt"

func main() {

	letters := []string{"a", "b", "c"}
	for i := 0; i < len(letters); i++ {
		fmt.Println(letters[i])
	}
	fmt.Println(" ")
	for _, letter := range letters {
		fmt.Println(letter)
	}
}
