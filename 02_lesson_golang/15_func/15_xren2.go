package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Print("Enter a: ")
	fmt.Scan(&a)
	fmt.Print("Enter b: ")
	fmt.Scan(&b)
	fmt.Println(a+b, " = a + b")
}
