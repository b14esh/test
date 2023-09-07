package main

import "fmt"

func za(i int) {
	for x := 1; x <= i; x++ {
		fmt.Printf("%d \n ", x)
	}
}

func main() {
	za(100)
}
