package main

import "fmt"

func main() {

	var (
		name     string
		lastname string
		age      int
	)

	fmt.Println("Введите ваше имя, фамелию, возраст")
	fmt.Scan(&name, &lastname, &age)

	fmt.Println("Имя:", name, "\nФамилия:", lastname, "\nВозраст:", age)
}
