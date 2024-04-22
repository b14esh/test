package main

import "fmt"


func main() {
	a := "Hello world"
	b := "No Hello world"
	c := "Helo helo world"

	age := 200

	if age < 50 {
		fmt.Println(a)
	} else if age == 100 {
		fmt.Println(b)
	} else if (age > 190) && (age <= 200) {
		var booble = age -5
		fmt.Println("You have", booble, "Yah", "and you age =", age)
	} else {
		fmt.Println(c)
	}
}
