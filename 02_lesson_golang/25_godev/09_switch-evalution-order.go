package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	case today + 3:
		fmt.Println("3?")
	case today + 4:
		fmt.Println("4?")
	case today + 5:
		fmt.Println("5?")
	case today + 6:
		fmt.Println("6?")
	case today + 7:
		fmt.Println("7?")
	default:
		fmt.Println("Too far away.")
	}
}
