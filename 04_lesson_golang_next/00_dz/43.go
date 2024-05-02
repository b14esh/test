// /Проверяем можно ли разделить арбуз по палам.
// Каждая часть тоже должна делиться по палам.
// 8 YES
package main

import (
	"fmt"
)

func main() {

	var result int
	var proverka int
	fmt.Scan(&result)
	if result < 1 || result > 100 {
		fmt.Println("NO")
	}
	if result%2 == 0 {
		proverka = result / 2
		if proverka%2 == 0 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

	} else {
		fmt.Println("NO")
	}
}
