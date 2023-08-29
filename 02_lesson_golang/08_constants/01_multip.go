package main

import "fmt"

const (
	pi   = 3.14
	lang = "GO"
	x    = 100
	b    = 200
	c    = 50
)

func main() {
	fmt.Println(pi)
	fmt.Println(lang)
	fmt.Println("x + b + c =", c+b+x)
}
