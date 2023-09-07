package main

import "fmt"

func xox(a int, b int) {
	fmt.Printf("%d - %b\n", a, b)
}

func main() {
	c := 10
	y := 1
	for x := 1 ; x <= c ;  x++{
	    xox(x, y)
		y++
	}
}
