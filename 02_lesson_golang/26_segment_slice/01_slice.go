package main

import "fmt"

func main() {
	pr := make([]int, 5)
	pr[0] = 99
	pr[4] = 200
	fmt.Println("pr[4] - pr[0] = ", pr[4]-pr[0])

}
