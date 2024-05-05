package main

import "fmt"

func TypeFinder(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Println("This is string")
	case int:
		fmt.Println("this is int")
	default:
		fmt.Println("Unknow type")
	}
}

func main() {
	TypeFinder(10)
	TypeFinder("Hello")
	TypeFinder(nil)
}
