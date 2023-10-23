package main

import "fmt"

func main() {
	var value int = 2         //Создает значение
	var pointer *int = &value //Получает указатель на это значение.
	fmt.Println(pointer)      //юрприз! Выводится указа-тель, а не значение!
}
