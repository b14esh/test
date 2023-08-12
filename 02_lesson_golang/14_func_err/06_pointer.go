package main

import "fmt"

func pointer(x *int) {
	*x = 2
}

func main() {
	x := 0
	pointer(&x)
	fmt.Println(x)
}
