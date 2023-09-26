package main

import (
	"fmt"
)

func f() int {
	return 100
}
func main() {
	i := 100
	switch i {
	case 0:
		fmt.Println(i+100)
	case f():
		fmt.Println(i-50)
	default:
		fmt.Println("Hi go", i+30)
	}
}
