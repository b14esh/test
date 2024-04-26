package main

import "fmt"

func main() {

	for {
		var price int
		fmt.Scan(&price)

		switch price {
		case 100:
			fmt.Println("A case")
		case 110:
			fmt.Println("B case")
		case 120:
			fmt.Println("C case")
		case 130:
			fmt.Println("E case")
		case 999:
			fmt.Println("go return")
			return
		default:
			// Отрабатывает только в том случае если только ни один з выше не сработал
			fmt.Println("No case")
		}

	}
}
