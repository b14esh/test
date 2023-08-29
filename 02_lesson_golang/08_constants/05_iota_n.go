package main

import "fmt"

const (
	_ = iota      // 0
	b = iota * 10 // 1 * 10
	c = iota * 10 // 2 * 20
	e = iota * 10 // 3 * 30
	d = iota * 10 // 4 * 30
)

func main() {
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(e)
	fmt.Println(d)
}
