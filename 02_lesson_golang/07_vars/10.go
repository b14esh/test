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

	fmt.Printf("%v \t\t= %T \n", a, a)
	fmt.Printf("%v \t\t= %T \n", b, b)
	fmt.Printf("%v \t\t= %T \n", c, c)
	fmt.Printf("%v \t\t= %T \n", d, d)
	fmt.Printf("%v \t\t= %T \n", e, e)
	fmt.Printf("%v \t= %T \n", f, f)
	fmt.Printf("%v \t\t= %T \n", g, g)
}
