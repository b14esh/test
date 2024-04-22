package main 

import "fmt"

func main() {
	//  Переключатель
	var age = 9
	switch age {
	case 5: fmt.Println("You age 5")
	case 15: fmt.Println("You age 15")
	case 20: fmt.Println("You age 20")
	case 30: fmt.Println("You age 30")
	case 10: fmt.Println("You age 10")
	default: fmt.Println("I dont know")
	}	
}
