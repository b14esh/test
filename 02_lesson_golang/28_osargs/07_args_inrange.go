package main

import "fmt"

func inRange(min float64, max float64, numbers ...float64) []float64 {
	var result []float64
	for _, number := range numbers {
		if number >= min && number <= max {
			result = append(result, number)
		}
	}
	return result
}

func main() {
	fmt.Println(inRange(1, 100, -12.5, 4.2, 0, 50, 100.3)) // Поиск аргументов от 1 до 100 (задают первые два числа
	fmt.Println(inRange(-10, 10, 4.1, 12, -12, -5.2))      //Поиск аргументов от -10 до 10 (задают первые два числа)
}
