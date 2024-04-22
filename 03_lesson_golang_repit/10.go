package main

import "fmt"

func main() {
// массивы
	var arr[3] int
	arr[0] = 45
	arr[1] = 97
	arr[2] = 76
	fmt.Println(arr[1], arr[2])


	nums := [3]float64 {1.2, 3.3, 99.2}

	for i, value := range nums {
		fmt.Println(value, i)
	}
}
