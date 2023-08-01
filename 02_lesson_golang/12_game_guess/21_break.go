package main

import "fmt"

func main() {

	for x := 1; x <= 3; x++ {
		fmt.Println("before break")
		break
		fmt.Println("after break")
	}
	fmt.Println("after loop")

}
