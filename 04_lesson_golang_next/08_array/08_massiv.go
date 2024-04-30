package main

import "fmt"

func main() {

	//Перебор массива
	//Итерирование по массиву
	floatArr := [...]float64{10.1, 15.4, 44.44, 12.5}
	for i := 0; i < len(floatArr); i++ {
		fmt.Println(floatArr[i])
	}

	for i := 0; i < len(floatArr); i++ {
		fmt.Printf("%d element of arr if %.2f \n", i, floatArr[i])
	}
}
