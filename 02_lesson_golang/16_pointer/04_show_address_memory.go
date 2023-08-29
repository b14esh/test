package main

import "fmt"

func main() {

	a := 100

	fmt.Println("a =", a)
	fmt.Println("a is memory address =", &a)
	fmt.Printf("a is decimal = %d \n", &a)
	fmt.Printf("a is bin = %b \n", &a)
	fmt.Printf("a is hex = %x \n", &a)
}
