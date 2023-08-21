package main

import "fmt"

func znak() {
	a := 4
	x := &a
}

func main() {

	x := *x
	fmt.Println(x + x)

}
