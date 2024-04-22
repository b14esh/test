package main

import "fmt"

func main() {
	//maps карты
	webS := make(map[string]float64)

	webS["xxxx"] = 0.7
	webS["yyyy"] = 0.99

	fmt.Println(webS["xxxx"])
}
