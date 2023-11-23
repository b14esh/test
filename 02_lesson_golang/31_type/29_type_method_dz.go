package main

import "fmt"

type Population int //Обьявляем тип Population с базовым типом "int"

func main() {
	var population Population
	population = Population(572) //Целое число преобразуется в значение Population.
	fmt.Println("Sleepy Creek County population:", population)
	fmt.Println("Congratulations, Kebin and Anna! It's a girl!")
	population += 1
	fmt.Println("Sleepy Creek County population:", population)
}
