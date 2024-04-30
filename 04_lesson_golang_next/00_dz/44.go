// Пирамида
package main

import "fmt"

/*//раз
func main() {
	rows := 4
	lastChar := 'A'

	for i := 1; i <= rows; i++ {
		for j := 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}

		for k := 1; k <= 2*i-1; k++ {
			fmt.Printf("%c", lastChar)
		}

		lastChar++
		fmt.Println()
	}
}
*/
/*
//два
func main() {
	rows := 4
	symbol := '*'

	for i := 1; i <= rows; i++ {
		for j := 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}

		for k := 1; k <= 2*i-1; k++ {
			fmt.Printf("%c", symbol)
		}
		fmt.Println()
	}
}
*/
//три
/*
func main() {

	var inter int
	//fmt.Scan(&inter)
	inter = 3

	for a := 0; a < inter; a++ {
		for b := inter; b > a; b -= 1 {
			fmt.Print(" ")
		}
		for c := 0; a > c; c++ {
			fmt.Print("@")
		}
		for d := 0; d < a; d++ {
			fmt.Print(" ")
		}
		fmt.Println()
	}

}

*/

func main() {

	var rows int
	fmt.Scan(&rows)

	//rows := 3
	symbol := '@'

	for i := 1; i <= rows; i++ {
		for j := 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}

		for k := 1; k <= 2*i-1; k++ {
			fmt.Printf("%c", symbol)
		}
		fmt.Println()
	}
}
