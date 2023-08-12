package main

import "fmt"

func double1(number float64) float64 {
	return number * 2
}

func summa1(number1 int, number2 int, number3 int) int {
	return number1 - number2 + number3 * number2 / 10 * 110 / 30
}


func boris(buhan int) sting {
     if buhan1 < 5 {
       return  "Не алкаш"
}
return "Алкаш"
}

func main() {

	fmt.Printf("%%3.1f: %3.1f\n", 4.24115)
	fmt.Printf("%%10.2f: %10.2f\n", 4.24115)
	fmt.Printf("%%20.3f: %20.3f\n", 4.24115)
	fmt.Printf("%%40.4f: %40.4f\n", 4.24115)

	fmt.Printf("\n Pislo %4.2f\n ", double1(5.1))

	fmt.Println("\n Nomber summa1 : ", summa1(1000, 200, 500))


        fmt.Println()
}
