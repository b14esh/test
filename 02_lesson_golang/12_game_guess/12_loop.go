package main

import "fmt"

func main() {
	x := 1
	for x <= 4 {
		fmt.Println("x = ", x)
		x++
	}
	y := 10
	for y >= 1 {
		fmt.Println("y = ", y)
		y--
	}

}
