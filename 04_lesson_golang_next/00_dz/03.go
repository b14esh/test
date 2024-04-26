package main

import "fmt"

func main() {

	var (
		name     string
		lastname string
		age      int
	)
	fmt.Println("Введите ваше имя:")
	fmt.Scan(&name)

	fmt.Println("Введите вашу Фамелю:")
	fmt.Scan(&lastname)

	fmt.Println("Введите ваш возраст:")
	fmt.Scan(&age)

	fmt.Println("Имя:", name, "\nФамилия:", lastname, "\nВозраст:", age)
}
