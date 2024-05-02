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

//3. Функция как параметр в другой функции
func reciveFuncAndreturnValue(f func(a, b int) int) int {
	var intVarA, intVarB int
	intVarA = 100
	intVarB = 200
	return f(intVarA, intVarB)
}

func add(a, b int) int {
	return a + b
}

func main() {
	var command string
	//command = "addition"
	command = "substration"
	res := calcAndReurnValidFunc(command)
	fmt.Println("res with command", command, "value", res(10, 20))
	fmt.Println(res(30, 40))

	//4. Тип функции
	fmt.Printf("Type of func is %T\n", res)
	fmt.Printf("Type of calcAndReurnValidFunc is %T\n", calcAndReurnValidFunc)
	//5. Тип функции го определяется входными так и выходными параметрами

	fmt.Println("recivefuncandReturnValue(add)", reciveFuncAndreturnValue(add))
	fmt.Println(reciveFuncAndreturnValue(func(a, b int) int {
		return a*a + b*b + 2*a*b
	}))

}
