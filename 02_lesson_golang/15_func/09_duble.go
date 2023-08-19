package main

import "fmt"

func main() {
	amount := 6
	double(amount)      //Функция передает аргумент
	fmt.Println(amount) //Выводится исходное значение
}

func double(number int) { //Параметру присваивается копия аргумента
	number *= 2 // Измен
}
