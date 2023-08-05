package main

import "fmt"

func xxx(x string) {
	fmt.Println(x)
}

var x1 int
var x2 int
var x3 int

func calc1(a int, b int, c int) {
	x1 = a + b
	x2 = c - b
	x3 = a + b + c
	return
}

func main() {
	for i := 1; i < 10; i++ {
		xxx("string")
	}

	calc1(10, 20, 30)
	fmt.Println(x1)
	fmt.Println(x2)
	fmt.Println(x3)

}
