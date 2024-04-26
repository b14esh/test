package main

import "fmt"

func main() {
	var (
		age  int
		name string
	)

	fmt.Println("Input you age:")
	fmt.Scan(&age)
	fmt.Println("Input you name:")
	fmt.Scan(&name)

	//fmt.Scan(&age, &name)
	// тогда можно будет записывать вот так вот "16 Name" или "Name 16"

	fmt.Printf("My name is: %s \nMy age is: %d", name, age)
}
