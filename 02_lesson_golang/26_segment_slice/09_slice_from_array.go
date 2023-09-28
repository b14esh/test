package main

import "fmt"

func main() {
	underlyingArray := [5]string{"a", "b", "c", "d", "e"}
	slice4 := underlyingArray[:3]
	fmt.Println(slice4)
}
