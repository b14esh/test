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
	/*
		a := 35231
		sa := strconv.Itoa(a)
		si := ""
		xsi := 0
		box := []rune(sa)
		result := []int{}

		for i := 0; i < len(box); i++ {
			fmt.Printf("%c\n", box[i])
		}

		fmt.Println()

		for i := len(box) - 1; i >= 0; i-- {
			fmt.Printf("%c", box[i])
			si = fmt.Sprintf("%c", box[i])
			xsi, _ = strconv.Atoi(si)
			result = append(result, xsi)

		}

		fmt.Println()
		fmt.Println(result)
	*/

	a := 35231
	sa := strconv.Itoa(a)
	si := ""
	xsi := 0
	box := []rune(sa)
	result := []int{}

	for i := len(box) - 1; i >= 0; i-- {
		si = fmt.Sprintf("%c", box[i])
		xsi, _ = strconv.Atoi(si)
		result = append(result, xsi)

	}
	fmt.Println(result)

}

/*
package kata

import (
        "fmt"
        "strconv"
)

func Digitize(n int) []int {
  a := n
	sa := strconv.Itoa(a)
	si := ""
	xsi := 0
	box := []rune(sa)
	result := []int{}

	for i := len(box) - 1; i >= 0; i-- {
		si = fmt.Sprintf("%c", box[i])
		xsi, _ = strconv.Atoi(si)
		result = append(result, xsi)

	}
	return result
}

*/

/*

package kata

func Digitize(n int) []int {
	var ret []int
	for {
		ret = append(ret, n%10)
		n /= 10
		if n == 0 {
			return ret
		}
	}
}

*/
