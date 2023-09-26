package main

import "fmt"

func main() {
	for i := 15; i > 1; i-- {
		switch {
		case i < 5:
			fmt.Println("i < 4")
		case i < 7:
			fmt.Println("i < 7")
		case i < 10:
			fmt.Println("i < 10")
		default:
			fmt.Println("I'm tired!")
		}
	}
}
