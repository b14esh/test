package main

import "fmt"

func main() {
	grades := map[string]float64{"Alma": 74.2, "Rohit": 86.5, "Carl": 59.7}
	for name := range grades {
		fmt.Println(name)
	}
}
