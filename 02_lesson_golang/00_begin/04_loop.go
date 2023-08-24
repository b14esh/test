package main

import "fmt"

func main() {
	for i := 1; i < 256; i++ {
		fmt.Printf("%d \t %b \t %x \n", i, i, i)
	}
	fmt.Println()
}
