package main

import "fmt"

func main() {

	var (
		drink   string
		meal    string
		subMeal string
		time    string
	)

	fmt.Println("Enter please: drink, meal, subMeal, time")
	fmt.Println("Example: tea chicken fries 13:00")
	fmt.Println("Example: coffee beef pie 12:30")

	fmt.Scan(&drink, &meal, &subMeal, &time)

	fmt.Printf("I wanna some '%s', my favorite meal -'%s'. Give me also '%s'. I will come soon... '%s'", drink, meal, subMeal, time)

}
