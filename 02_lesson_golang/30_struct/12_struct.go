package main

import "fmt"

type subscriber struct {
	name   string
	rate   float64
	active bool
}

//Объявляется один параметр
//с типом «subscriber».
func printInfo(s subscriber) {
	fmt.Println("Name:", s.name)
	fmt.Println("Monthly rate:", s.rate)
	fmt.Println("Active?", s.active)
}

func defaultSubscriber(name string) subscriber { //Возвращает значение «subscriber».
	var s subscriber //Создается новое значение «subscriber».
	s.name = name
	s.rate = 5.99
	s.active = true
	return s //Возвращает «subscriber».
}

func main() {
	subscriber1 := defaultSubscriber("Aman Singh") //Создает subscriber с задан-ным значением name.
	subscriber1.rate = 4.99                        //Использует заданное значение rate.
	printInfo(subscriber1)                         //Вывод значений полей

	subscriber2 := defaultSubscriber("Beth Ryan") //Создает subscriber с задан-ным значением name.
	printInfo(subscriber2)                        //Вывод значений полей.
}
