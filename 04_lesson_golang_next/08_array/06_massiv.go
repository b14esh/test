package main

import "fmt"

func main() {

	//Массив - это набор ЗНАЧЕНИЙ. То есть при всех манипуляциях - массив копируется (жестко, на уровне компилятора)
	first := [...]int{1, 2, 3}
	second := first
	second[0] = 100000

	fmt.Println("First arr:", first)
	fmt.Println("Second arr:", second)
}
