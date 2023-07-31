package main

import "fmt"

func main() {
        var y int
	for x := 1; x <= 3; x++ {
		y = x + 1
		fmt.Println(y)
	}
	fmt.Println(y)
}
