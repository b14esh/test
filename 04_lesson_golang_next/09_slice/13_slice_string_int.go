package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	// The int slice we are converting to a string.
	values := []int{10, 200, 3000}
	valuesText := []string{}

	// Create a string slice using strconv.Itoa.
	// ... Append strings to it.
	for i := range values {
		number := values[i]
		text := strconv.Itoa(number)
		valuesText = append(valuesText, text)
	}

	// Join our string slice.
	result := strings.Join(valuesText, "+")
	fmt.Println(result)
}
