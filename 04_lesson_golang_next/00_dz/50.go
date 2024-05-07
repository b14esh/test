package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	box := []string{}
	number := 5
	sx := ""
	si := ""
	sn := strconv.Itoa(number)
	//result := []string{}

	var x int
	for i := 1; i <= 10; i++ {
		x = i * number
		si = strconv.Itoa(i)
		sx = strconv.Itoa(x)
		box = append(box, si)
		box = append(box, " * ")
		box = append(box, sn)
		box = append(box, " = ")
		box = append(box, sx)
		if i != 10 {
			box = append(box, "\n")
		}
	}

	fmt.Println(box)
	zx := "\"" + strings.Join(box, "") + "\""

	fmt.Println(zx)

}

/*
func MultiTable(number int) string {
	superstring := ""
	for i := 1; i < 11; i++ {
		superstring += fmt.Sprintf("%d * %d = %d\n", i, number, number * i)
	}
	return strings.TrimRight(superstring, "\n")
}

*/

/*
func MultiTable(number int) (result string) {
	for i := 1; i <= 10; i++ {
		result += fmt.Sprintf("%v * %v = %v", i, number, i*number)
		if i != 10 { result += "\n"}
	}
	return
}
*/

/*
func MultiTable(number int) string {
  MultiString := ""
  for i := 1; i <= 10; i = i + 1 {
    MultiString = MultiString + strconv.Itoa(i) + " * " + strconv.Itoa(number) + " = " + strconv.Itoa(i*number) + "\n"
  }
  return strings.TrimSuffix(MultiString, "\n")
}
*/

/*
import (
	"fmt"
)

func MultiTable(number int) string {
  var res string
	for i := 1; i <= 10; i++ {
		res += fmt.Sprintf("%d * %d = %d", i, number, i*number)
		if i < 10 {
			res += "\n"
		}
	}
	return res
}

*/
