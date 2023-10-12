package main

import "fmt"

type subscriber struct {
	name   string
	rate   float64
	active bool
}

func applyDiscount(s subscriber) {
	s.rate = 4.99 //Получает параметр «subscriber».
}
func main() {
	var s subscriber
	applyDiscount(s)    //Пытается присвоить полю «rate» структуры «subscriber» значение 4.99
	fmt.Println(s.rate) //Поле остается равным 0!
}
