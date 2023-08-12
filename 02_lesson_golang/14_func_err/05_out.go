package main

import "fmt"

func one() {
	fmt.Println("1")
}

func two() {
	fmt.Println("2")
}

//откладывание

func main() {
	defer two()
	one()
}
