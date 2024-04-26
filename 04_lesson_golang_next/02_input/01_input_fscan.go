package main

import (
	"fmt"
	"os"
)

func main() {

	var (
		age  int
		name string
	)

	fmt.Println("Enter age and name.")
	fmt.Scan(&age, &name)
	fmt.Printf("My name is: %s \nMy age is: %d", name, age)

	//Для ручного использования потока ввода
	fmt.Println("\nEnter new age:")
	fmt.Fscan(os.Stdin, &age)
	fmt.Println("New age:", age)

}
