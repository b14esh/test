package main

import "fmt"

func main() {

	//Кейсы с выражениями
	var age int
	fmt.Scan(&age)
	switch {
	case age <= 18:
		fmt.Println("you yong", age)
	case age > 18 && age <= 30:
		fmt.Println("ABUAU")
	case age > 30 && age <= 100:
		fmt.Println("To old")
	default:
		fmt.Println("Who are you?")
	}
}
