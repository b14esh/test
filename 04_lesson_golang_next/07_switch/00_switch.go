package main

import "fmt"

func main() {

	//Switch !

	//Зачем? Представим что у нас простыня из ифов:
	//  price := 120
	// if price == 100 {
	//	fmt.Println("Firs case")
	//  } else if price == 110 {
	//	fmt.Println("Second case")
	//  } else if price == 120 {
	//	mft.Println("ALA case")
	//  } ....

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
	}

}
