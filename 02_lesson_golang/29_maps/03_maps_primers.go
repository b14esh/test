package main

import "fmt"

func main() {

	fmt.Println("\nPrint map numbers:")
	ranks := make(map[string]int)
	ranks["gold"] = 1
	ranks["silver"] = 2
	ranks["bronze"] = 3
	fmt.Println(ranks["bronze"])
	fmt.Println(ranks["gold"])

	fmt.Println("\nPrint map letters:")
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["Li"] = "Lithium"
	fmt.Println(elements["Li"])
	fmt.Println(elements["H"])

	fmt.Println("\nPrint map bool:")
	isPrime := make(map[int]bool)
	isPrime[4] = false
	isPrime[7] = true
	fmt.Println(isPrime[4])
	fmt.Println(isPrime[7])
}
