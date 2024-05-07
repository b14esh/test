//Convert number to reversed array of digits
//Given a random non-negative number, you have to return the digits of this number within an array in reverse order.

//Example(Input => Output):
//35231 => [1,3,2,5,3]
//0 => [0]

/*
func Digitize(n int) []int {
	// your code here
	return nil
  }
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := 35231
	sa := strconv.Itoa(a)
	box := []rune(sa)

	for i := 0; i < len(box); i++ {
		fmt.Printf("%c\n", box[i])
	}

	fmt.Println()

	for i := len(box) - 1; i >= 0; i-- {
		fmt.Printf("%c\n", box[i])
	}

}
