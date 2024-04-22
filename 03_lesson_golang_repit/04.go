package main

import "fmt"

func main() {
	var num int = 100
	var fff  float64 = 200.123
        var str string = "Ehlo man"
	fmt.Printf(" Тут будет int = %v \n Тут будет float = %f \n Тут будет string %s \n", num, fff, str)
	
	fmt.Println("У нее есть специальный символ %T, что выводит тип переменной.")
	fmt.Printf(" Тут будет int = %T \n Тут будет float = %T \n Тут будет string %T \n", num, fff, str)

}
