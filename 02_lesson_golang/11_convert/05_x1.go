package main

import "fmt"

func main() {
	a := 1
        //a := 2
        fmt.Println(a)
	b, a := 2, 3
        fmt.Println(b, a)
	a, c := 4, 5
	fmt.Println(a, b, c)
}
