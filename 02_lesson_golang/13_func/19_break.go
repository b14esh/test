package main

import "fmt"

func paintNeeded(width float64, height float64) float64 {
//func paintNeeded(width float64, height float64) /* float64*/ {
	area := width * height
	return area / 10.0
//	return area / 10.0
//	fmt.Println(area / 10.0)
//	return int(area / 10.0)
}
func main() {
	var amount, total float64
	amount = paintNeeded(4.2, 3.0)
	fmt.Printf("%0.2f liters needed\n", amount)
	total += amount
	amount = paintNeeded(5.2, 3.5)
	fmt.Printf("%0.2f liters\n", total)
	total += amount
	fmt.Printf("Total: %0.2f litters\n", total)
}
