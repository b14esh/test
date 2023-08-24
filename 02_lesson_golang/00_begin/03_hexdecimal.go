package main

import "fmt"

func main() {
	fmt.Printf("%d - %b - %x \n", 13, 13, 13)
	fmt.Printf("%d -%b - %#x \n", 13, 13, 13)
	fmt.Printf("%d - %b - %#X \n", 13, 13, 13)
	fmt.Printf("%d \t %b \t %#X \n", 13, 13, 13)
}
