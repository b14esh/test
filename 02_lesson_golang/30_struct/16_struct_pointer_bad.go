package main

import "fmt"

func main() {
	var value int = 2         // Создаем значение.
	var pointer *int = &value // Получаем указатель на это значение.
	fmt.Println(pointer)      // Сюрприз! ВЫводится указатель, а не значение!
}
