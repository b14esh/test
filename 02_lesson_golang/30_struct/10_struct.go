package main

import "fmt"

type part struct {
	description string
	count       int
}

func showInfo(p part) { //Объявление одного параметра с типом «part».
	fmt.Println("Description:", p.description)
	fmt.Println("Count:", p.count)
}

func main() {
	var bolts part //Создается значение «part».
	bolts.description = "Hex bolts"
	bolts.count = 24
	showInfo(bolts) //Тип «part» пере-дается функции
}
