package main

import (
	"fmt"
)

func main() {

	var result int
	fmt.Scan(&result)
	if result < 1 || result > 100 {
		fmt.Println("NO")
	}
	if result%2 == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
