package main

import "fmt"

func main() {

	//Свитч с множеством вариантов
	var ageGroup string = "wx -w1 -a -h" //Возрасные группы: "a", "b", "c", "d", "e"
	switch ageGroup {
	case "wx -w1 -a -h":
		fmt.Println("AgeGroup 10-40")
	case "d", "e":
		fmt.Println("AgeGroup 50-70")
	default:
		fmt.Println("You a too yong/old")

	}

}
