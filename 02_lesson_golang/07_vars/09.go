package main

import "fmt"

func main() {
	a := 15
	b := "hulio"
	c := 9.13
	d := true
	e := "ehlo"
	f := `chto take`
	g := 'M'

	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
	fmt.Printf("%T \n", e)
	fmt.Printf("%T \n", f)
	fmt.Printf("%T \n", g)
}
