package main

import "fmt"

func main() {

	var price int
	fmt.Scan(&price)
	// В switch case запрещенны дубли
	switch price {
	case 100:
		fmt.Println("A case")
	case 110:
		fmt.Println("B case")
	case 120:
		fmt.Println("C case")
	case 130:
		fmt.Println("E case")
	default:
		// Отрабатывает только в том случае если только ни один з выше не сработал
		fmt.Println("No case")

	}
}
