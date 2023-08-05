package main

import "fmt"

var metersPerLiter float64 // Если переменная объявляется на уровне пакета...
func paintNeeded(width, height float64) float64 {
	area := width * height
	return area / metersPerLiter //здесь она остается в области видимости.
}
func main() {
	metersPerLiter = 10.0 //и здесь остается в области видимости.
	fmt.Printf("%.2f\n", paintNeeded(4.2, 3.0))
}
