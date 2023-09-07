package main

import "fmt"

func test1(a int) {
	fmt.Printf("%d \t %b \t %#X \n", a, a, a)
}

func main() {
	for a := 1; a <= 101; a++ {
		test1(a)
	}
}
