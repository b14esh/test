package main

import "fmt"

func main() {

	box := map[int]string{1: "a", 2: "b", 3: "c"}
	for k, x := range box {
		fmt.Println(k, "==", x)
	}
}
