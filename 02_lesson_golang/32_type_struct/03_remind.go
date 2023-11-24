package main

import "fmt"

type Date struct {
	Year  int
	Month int
	Day   int
}

func main() {
	date := Date{Year: 123, Month: -1, Day: 0}
	fmt.Println(date)
	fmt.Println(date.Day)
	fmt.Println(date.Year)
	fmt.Println(date.Month)
	date = Date{Year: -123, Month: 9999, Day: 50}
	fmt.Println(date)
	date = Date{Year: 666, Month: 0, Day: -666}
	fmt.Println(date)
}
