package main

import "fmt"

func main() {
	var count int = 12
	var suffix string = "minutes of bonus footage"
	var format string = "DVD"
	var languages = append([]string{}, "Español")
	fmt.Println(count, suffix, "on", format, languages)
}
