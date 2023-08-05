package main

import "fmt"

func status(grade float64) string {
	if grade < 60.0 {
		return "failing" // Если экзамен провален, немедленно вернуть.
	}
	return "passing" // Выполняется только в том случае, если grade >= 60.
}
func main() {
	fmt.Println(status(60.1))
	fmt.Println(status(59))
}
