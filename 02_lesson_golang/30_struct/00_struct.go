package main

import "fmt"

func main() {
	var myStruct struct {
		number float64
		world  string
		toggle bool
	}
	fmt.Printf("%#v\n", myStruct)
}
