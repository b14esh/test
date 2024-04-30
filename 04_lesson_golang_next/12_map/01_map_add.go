package main

import "fmt"

func main() {

	// Map , карта, это паразначений ключ и значение. Иницилизация пустой карты / мапы
	mapper := make(map[string]int)

	fmt.Println("Empty map:", mapper)

	//Добавление пар в существующую мапу
	mapper["Alice"] = 24
	mapper["Bob"] = 25
	fmt.Println("Mappert afte add", mapper)

}
