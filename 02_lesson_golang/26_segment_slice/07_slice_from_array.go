package main

import "fmt"

func main() {

	underlyingArray := [5]string{"a", "b", "c", "d", "e"}
	i, j := 1, 4
	slice2 := underlyingArray[i:j]
	fmt.Println(slice2)
	fmt.Println(slice2[1])

}
