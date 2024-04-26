package main

import "fmt"

func main() {
	fmt.Println("Hellow world!")

	//Класический условный оператор
	//if condition {
	//	//body
	//
	//}

	// Условный оператор с блоком else
	//if condition {

	//} else {

	//}

	var value int
	fmt.Println("Введите число:")
	fmt.Scan(&value)

	if value%2 == 0 { // если число делится без остатка то оно четное //  % - остаток от делени
		fmt.Println("The number", value, "is even") // четное число

	} else {
		fmt.Println("The number", value, "is odd") // не четное

	}

	//if condition1 {

	//} else if condition2 {

	//} else if ... {

	//} else {

	//}

}
