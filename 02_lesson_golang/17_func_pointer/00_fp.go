package main

import "fmt"

func createInt() *int {
	var myInt = 100
	return &myInt
}

func createPointer() *float64 {
	var myFloat = 98.5
	return &myFloat
}

func crS() *string {
	var xX = "hello all"
	return &xX
}


func main() {
	var myFloatPointer *float64 = createPointer()
	fmt.Println(*myFloatPointer)

	var x *int = createInt()
	fmt.Println(*x)

	var x1 *string = crS()
	fmt.Println(*x1)
}
