package main

import "fmt"

func main() {
	for x := 1; false; x++ {
		fmt.Println(x)
	}
}
