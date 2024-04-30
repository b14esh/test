package main

import "fmt"

func main() {

	//Перебор массива крассиво
	//Итерирование по массиву через оператор range
	floatArr := [...]float64{10.1, 15.4, 44.44, 12.5}
	//оператор range всегда возращает два элемента!

	var sum float64

	for id, val := range floatArr {
		fmt.Printf("%d element of arr is %.2f \n", id, val)
		sum += val
	}
	fmt.Println("Total sum i", sum)
}
