package main

import "fmt"

//1. Явная функция (в принципе любая функция Go) - является
// экземпляром 1-го уровня (функцию можно присваивать в переменную,
// ее можно передать в качестве  параметра и возращать из других функций)

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mult(a, b int) int {
	return a * b
}

func calcAndReurnValidFunc(command string, a int, b int) int {
	if command == "addition" {
		return func(a, b int) int { return a + b }(a, b)
	} else if command == "substration" {
		return func(a, b int) int { return a - b }(a, b)
	} else {
		return func(a, b int) int { return a * b }(a, b)
	}
}

func main() {

	var command string
	command = "addition"
	res := calcAndReurnValidFunc(command, 10, 20)
	fmt.Println("res with command", command, "value", res)
}
