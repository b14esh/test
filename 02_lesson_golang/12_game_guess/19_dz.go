package main

import "fmt"

func main() {

	fmt.Println("\nz1")
	for x := 1; x <= 3; x++ {
		fmt.Print(x)
	}

	fmt.Println("\n\nz2")
	for x := 3; x >= 1; x-- {
		fmt.Print(x)
	}

	fmt.Println("\n\nz3")
	for x := 2; x <= 3; x++ {
		fmt.Print(x)
	}

	fmt.Println("\n\nz4")
	for x := 1; x < 3; x++ {
		fmt.Print(x)
	}

	fmt.Println("\n\nz5")
	for x := 1; x <= 3; x += 2 {
		fmt.Print(x)
	}

	fmt.Println("\n\nz6")
	for x := 1; x >= 3; x++ {
		fmt.Print(x)
	}

}
