package main

import "fmt"

func main() {
	//Добавление элементов в срез
	myWorlds := []string{"one", "two", "three"}
	fmt.Println("myWords", myWorlds)
	myWorlds = append(myWorlds, "four")
	fmt.Println("myWords", myWorlds)
}
