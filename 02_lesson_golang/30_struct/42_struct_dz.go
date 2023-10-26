package main

import "fmt"

type student struct { //Определение типа структуры student
	name  string
	grade float64
}

func printInfo(s student) { //Опеделение функции, которая получает структуру student в параметре.
	fmt.Println("Name:", s.name)
	fmt.Printf("Grade: %0.1f\n", s.grade)
}

func main() {
	var s student
	s.name = "Alonzo Cole"
	s.grade = 92.3
	printInfo(s) //Функции передается структура
}
