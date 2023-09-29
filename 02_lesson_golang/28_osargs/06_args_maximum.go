package main

import (
	"fmt"
	"math"
)

func maximum(numbers ...float64) float64 { //Получаем любое количесво аргументов float64
	max := math.Inf(-1) //начинаем с очень низского значения
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return max
}

func main() {
	fmt.Println(maximum(71.8, 56.2, 89.5))
	fmt.Println(maximum(90.7, 89.8, 98.3, 99.2))
}
