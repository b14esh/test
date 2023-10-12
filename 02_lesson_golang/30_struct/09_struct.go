package main

import "fmt"

type subscriber struct { //Определение типа с име-нем «subscriber».
			 //Тип структуры используется для переменных как базовый для определения типа.
	name   string
	rate   float64
	active bool
}

func main() { //Объявление переменной типа «subscriber».
	var subscriber1 subscriber
	subscriber1.name = "Aman Singh"
	fmt.Println("Name:", subscriber1.name)
	var subscriber2 subscriber //Тип «subscriber» также использу-ется для второй переменной.
	subscriber2.name = "Beth Ryan"
	fmt.Println("Name:", subscriber2.name)
}
