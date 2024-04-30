// Изменение элементов среза

package main

import "fmt"

func main() {

	originArr := [...]int{30, 40, 50, 60, 79, 80}

	firstSlice := originArr[1:4] // Срез это набор ссылок нижележащего массива
	//firstSlice := originArr[:] // Создаст срез в весь массив
	for i, _ := range firstSlice {
		firstSlice[i]++
	}

	fmt.Println("OriginArr:", originArr)
	fmt.Println("FirstSlice:", firstSlice)

}
