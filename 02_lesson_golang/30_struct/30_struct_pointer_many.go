package main

import "fmt"

type subscriber struct {
	name   string
	rate   float64
	active bool
}

func printInfo(s *subscriber) { //Изменяется для передачи указателя.
	fmt.Println("Name:", s.name)
	fmt.Println("Monthly rate:", s.rate)
	fmt.Println("Active?", s.active)
}
func defaultSubscriber(name string) *subscriber { //Изменяется для воз-вращения указателя.
	var s subscriber
	s.name = name
	s.rate = 5.99
	s.active = true
	return &s //Возвращает указатель на структуру вместо самой структуры.
}
func applyDiscount(s *subscriber) {
	s.rate = 4.99
}
func main() {
	subscriber1 := defaultSubscriber("Aman Singh") //Это уже не структура, а указатель на структуру...
	applyDiscount(subscriber1) //Так как это структура, оператор & не нужен.
	printInfo(subscriber1)
	subscriber2 := defaultSubscriber("Beth Ryan")
	printInfo(subscriber2)
}
