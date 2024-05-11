/*
You will be given a list of strings.
You must sort it alphabetically (case-sensitive, and based on the ASCII values of the chars) and then return the first value.
The returned value must be a string, and have "***" between each of its letters.
You should not remove or add elements from/to the array.
*/
/*
   s = []string{"bitcoin", "take", "over", "the", "world", "maybe", "who", "knows", "perhaps"}
       Expect(TwoSort(s)).To(Equal("b***i***t***c***o***i***n"))
   s = []string{"turns", "out", "random", "test", "cases", "are", "easier", "than", "writing", "out", "basic", "ones"}
       Expect(TwoSort(s)).To(Equal("a***r***e"))


*/

package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	// sort int
	nums := []int{2, 1, 6, 5, 3, 4}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	fmt.Println(nums) // [1 2 3 4 5 6]

	//sort string
	s := []string{"bitcoin", "take", "over", "the", "world", "maybe", "who", "knows", "perhaps"}
	s = []string{"turns", "out", "random", "test", "cases", "are", "easier", "than", "writing", "out", "basic", "ones"}
	sort.Strings(s)
	fmt.Println(s) // [are basic cases easier ones out out random test than turns writing]

	// me need
	box := []string{}
	x := ""
	sm := ""

	sort.Strings(s)

	fmt.Println(s)

	sm = s[0]

	for i := 0; i < len(sm); i++ {
		x = string(sm[i]) + "***"
		box = append(box, x)
	}
	x = strings.Join(box, "")

	fmt.Println(x)

}

/*
package kata

import (
    "strings"
    "sort"
)

func TwoSort(arr []string) string {
  s := []string(arr)
  box := []string{}
	x := ""
	sm := ""
	sort.Strings(s)
	sm = s[0]
	for i := 0; i < len(sm); i++ {
    if i != len(sm)-1 {
		x = string(sm[i]) + "***"
    }else{
      x = string(sm[i])
    }
		box = append(box, x)
	}
	x = strings.Join(box, "")
	return x

}
*/

/*
package kata

import (
	"sort"
	"strings"
)

func TwoSort(arr []string) string {
	sort.Strings(arr)
	chars := strings.Split(arr[0], "")
	return strings.Join(chars, "***")
}

*/

/*
package kata
import (
  "strings"
  "sort"
)
func TwoSort(arr []string) string {
  sort.Strings(arr)
  return strings.Join(strings.Split(arr[0],""), "***")
}
*/
