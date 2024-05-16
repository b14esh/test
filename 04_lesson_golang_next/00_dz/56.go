/*
Finish the solution so that it sorts the passed in array of numbers.
If the function passes in an empty array or null/nil value then it should return an empty array.

For example:

solution(c(1, 2, 3, 10, 5)) # should return c(1, 2, 3, 5, 10)
solution(NULL)              # should return NULL


package kata
func SortNumbers(numbers []int) []int {

  return []int{} // your code here
}

Expect(Expect(SortNumbers([]int{1, 2, 10, 50, 5})).To(Equal([]int{1,2,5,10,50})))
    Expect(Expect(SortNumbers([]int{})).To(Equal([]int{})))

*/

package main

import (
	"fmt"
	"sort"
)

func SortNumbers(numbers []int) []int {
	switch {
	case numbers == nil:
		return []int{}
	default:
		sort.Slice(numbers, func(i, j int) bool {
			return numbers[i] < numbers[j]
		})
		return numbers
	}

}

func main() {
	Numbers := []int{1, 2, 10, 50, 5}
	fmt.Println(SortNumbers(Numbers))

}

/*
package kata
import (
	"sort"
)
func SortNumbers(numbers []int) []int {
	switch {
	case numbers == nil:
		return []int{}
	default:
		sort.Slice(numbers, func(i, j int) bool {
			return numbers[i] < numbers[j]
		})
		return numbers
	}
}
*/

/*
package kata

import "sort"

func SortNumbers(numbers []int) []int {
  sort.Ints(numbers)
  return numbers
}
*/
