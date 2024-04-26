package main

import "fmt"

func main() {
	//Простейший вывод. Println - это вывод аргумена + '\n'
	fmt.Println("Hello world")
	fmt.Println("Second line")
	//Функция Print - простой вывод аргумента
	fmt.Print("One")
	fmt.Print("Two")
	fmt.Print("Three")
	//Форматированый вывод Printf - стандартный вывод os.Stdout с флагами форматирования
	fmt.Printf("\n\nHello, my name is %s", "Bob")
	fmt.Printf("\t TAB \n \t \t double TAB")
}
