package main

import "fmt"

func sayHi() {
	fmt.Println("Hi!")
}

func main() {
	for x := 4; x > 1; x-- {
		sayHi()
	}
}
