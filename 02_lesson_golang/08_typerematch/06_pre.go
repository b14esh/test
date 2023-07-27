package main

import "fmt"

func main() {
	var length float64 = 1.2
	var width int = 2
	length = float64(width)
	fmt.Println(length)
}
