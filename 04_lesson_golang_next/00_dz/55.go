package main

import "fmt"

func main() {

	arr1 := []bool{
		true, true, true, false,
		true, true, true, true,
		true, false, true, false,
		true, false, false, true,
		true, true, true, true,
		false, false, true, true,
	}

	var x int
	numbers := arr1

	for i := 0; i < len(numbers); i++ {
		if numbers[i] == true {
			x = x + 1
		}
	}
	fmt.Println(x)

}

/*

package kata

func CountSheeps(numbers []bool) int {
  x := 0
  for i := 0; i < len(numbers); i++{
   if numbers[i] == true {
     x = x + 1
   }
  }
  return x
}

*/
