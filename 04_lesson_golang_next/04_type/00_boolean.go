package main

import "fmt"

func main() {
	//Boolean
	var firstBoolean bool
	// default false
	fmt.Println("Default bool:", firstBoolean)
	//Boolean operands
	aBoolean, bBoolean := true, false
	// логическое И, или как еще называют "логическое умножение"
	fmt.Println("AND:", aBoolean && bBoolean)
	// логическое ИЛИ, или логическое сложение
	fmt.Println("OR:", aBoolean || bBoolean)
	// логическое НЕ
	fmt.Println("NOT:", !aBoolean)
	// не будет работать как в питоне код не скопилируется
	// Err
	//# command-line-arguments
	//.\00_boolean.go:19:14: invalid operation: operator + not defined on aBoolean (variable of type bool)
	//fmt.Println(aBoolean + aBoolean)
	// Так как bool  имеет два значения "true или false"
}
