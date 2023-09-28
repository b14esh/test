package main

import "fmt"

func main() {
	underlyingArray := [5]string{"a", "b", "c", "d", "e"}
	slice3 := underlyingArray[2:5]
	fmt.Println(slice3)
}
