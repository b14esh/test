package main

import "fmt"

func main() {
	var value int = 2
	var pointer *int = &value
	fmt.Println(*pointer) // Выводим значение, на которое ссылается указатель.
}
