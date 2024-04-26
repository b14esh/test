package main

import "fmt"

func main() {

	var (
		a float64
		b float64
	)

	fmt.Println("Введи:\n 2.5 1.5")
	fmt.Scan(&a, &b)

	c := int(a + b)

	if num := c; num%2 == 0 {
		fmt.Println("ЧЁТНОЕ", num)
	} else {
		fmt.Println("НЕЧЁТНОЕ", num)
	}
}
