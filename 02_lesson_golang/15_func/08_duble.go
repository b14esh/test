package main

import "fmt"

func main() {
	amount := 6 //Аргумент передается функции
	double(amount)
}

func double(number int) { //В параметрах сохроняетя копия аргумента
	number *= 2
	fmt.Println(number) //Выводит удвоеное значение
}
