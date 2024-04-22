package main

import "fmt"

func main() {
	//Замыкания
	var num = 3
	mult := func() int {
		num *= 2
		return num
	}
	fmt.Println(mult())

}	
