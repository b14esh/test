package main

import "fmt"

func main() {
	//Добавление элементов в срез
	myWorlds := []string{"one", "two", "three"}
	fmt.Println("myWords", myWorlds)
	anotherSlice := []string{"four", "five", "six"}
	myWorlds = append(myWorlds, "four")

	//for _, val := range anotherSlice { // Не првильно но работать будет
	//	myWorlds = append(myWorlds, val)
	//}

	myWorlds = append(myWorlds, anotherSlice...)  // Правильный подход // Распаковка
	myWorlds = append(myWorlds, "seven", "eight") // По сути это делается

	fmt.Println("myWords", myWorlds)
}
