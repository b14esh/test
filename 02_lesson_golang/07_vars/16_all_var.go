package main

import "fmt"

var a = "lolo lolo a"                        //package scope
var b, c string = "store in b", "store in c" //package scope
var d string                                 //package scope

func main() {

	d = " strory in d"
	var e = 100
	f := 200
	g := "story in g"
	h, i := "story in h", "story in i"
	j, k, l, m := 44.7, true, false, 'm' // single quotes
	n := "n"                             // double quotes
	o := `o`                             //back ticks

	fmt.Println("a -", a)
	fmt.Println("b -", b)
	fmt.Println("c -", c)
	fmt.Println("d -", d)
	fmt.Println("e -", e)
	fmt.Println("f -", f)
	fmt.Println("g -", g)
	fmt.Println("h -", h)
	fmt.Println("i -", i)
	fmt.Println("j -", j)
	fmt.Println("k -", k)
	fmt.Println("l -", l)
	fmt.Println("m -", m)
	fmt.Println("n -", n)
	fmt.Println("o -", o)
}
