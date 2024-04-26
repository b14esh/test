package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hellow world!")

	//if condition1 {

	//} else if condition2 {

	//} else if ... {

	//} else {

	//}

	var color string

	fmt.Println("Введите цвет:")
	fmt.Scan(&color)

	if strings.Compare(color, "green") == 0 {
		fmt.Println("Color is green")
	} else if strings.Compare(color, "red") == 0 {
		fmt.Println("Color is red")
	} else {
		fmt.Println("Unknow color")
	}
}
