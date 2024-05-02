package main

import "fmt"

//1. Явная функция (в принципе любая функция Go) - является
// экземпляром 1-го уровня (функцию можно присваивать в переменную,
// ее можно передать в качестве  параметра и возращать из других функций)

//2. Возврат функции в качествек значения:
func calcAndReurnValidFunc(command string) func(a, b int) int {
	if command == "addition" {
		return func(a, b int) int { return a + b }
	} else if command == "substration" {
		return func(a, b int) int { return a - b }
	} else {
		return func(a, b int) int { return a * b }
	}
}

func main() {

	var command string

	//command = "addition"
	command = "substration"
	res := calcAndReurnValidFunc(command)
	fmt.Println("res with command", command, "value", res(10, 20))

	fmt.Println(res(30, 40))
}
