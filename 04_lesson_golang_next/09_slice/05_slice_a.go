// Улетающие ссылки в slice
package main

import "fmt"

func main() {

	numArr := [2]int{1, 2}
	numSlice := numArr[:]

	numSlice = append(numSlice, 3) // В этот момент nulSlice больше не ссылается на numArr
	numSlice[0] = 10

	fmt.Println(numArr)
	fmt.Println(numSlice)
}
