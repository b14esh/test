package main

import "fmt"

func main() {
	//Добавление элементов в срез
	myWorlds := []string{"one", "two", "three"}
	fmt.Println("myWords", myWorlds)
	anotherSlice := []string{"four", "five", "six"}
	myWorlds = append(myWorlds, "four")

	for _, val := range anotherSlice {
		myWorlds = append(myWorlds, val)
	}

	fmt.Println("myWords", myWorlds)
}
