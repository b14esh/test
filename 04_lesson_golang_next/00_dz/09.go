package main

import "fmt"

func main() {
	//Конкатенатор

	var (
		a string
		b string
		c string
	)

	fmt.Println("Введите три числа, \nнапример 2 4 6 \nили 1 1 2")
	fmt.Scan(&a, &b, &c)

	fmt.Println(c + b + a)

}
