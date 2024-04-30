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

	// Один массив и два произвольных среза
	fSlice := originArr[:]
	sSlice := originArr[2:5]

	fmt.Println("Befor modificatin: Arr:", originArr, "fSlice", fSlice, "sSlice", sSlice)
	fSlice[3]++ //Увеличим на 1
	sSlice[1]++ //Увеличим на 1
	fmt.Println("After modificatin: Arr:", originArr, "fSlice", fSlice, "sSlice", sSlice)

	// Срез как встроенный тип
	//type slice struct {
	//	Lenght int
	//	?????? int
	//	ZeroElement *byte
	//}

}
