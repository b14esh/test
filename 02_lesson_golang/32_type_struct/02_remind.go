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
}
