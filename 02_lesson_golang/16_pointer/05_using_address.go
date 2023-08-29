package main

import "fmt"

const metersToYards = 1.0961

func main() {
	var meters float64
	fmt.Printf("Enter meters: ")
	fmt.Scan(&meters)
	yards := meters * metersToYards
	fmt.Println(meters, " meters is ", yards, " yards.")
}
