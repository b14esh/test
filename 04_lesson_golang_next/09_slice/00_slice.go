package main

import "fmt"

func main() {
	// Слайсы (они же - срезы)
	// Слайс - это динамическая обьвязка над массивом.
	starArr := [4]int{10, 20, 30, 40}
	//Создадим слайз на основе существующего массива starArr
	var startSlice []int = starArr[:2] // Слайз иницилизируется пустыми квадратными скобками
	fmt.Println("slice[0:2]", startSlice)

}
