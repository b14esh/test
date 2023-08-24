package main

import "fmt"

func main() {
	fmt.Println("ENG")
	for i := 32; i < 128; i++ {
		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
	}
	fmt.Println()
}
