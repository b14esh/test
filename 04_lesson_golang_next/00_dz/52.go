/*
Is the string uppercase?
Task
Create a method to see whether the string is ALL CAPS.

Examples (input -> output)
"c" -> False
"C" -> True
"hello I AM DONALD" -> False
"HELLO I AM DONALD" -> True
"ACSKLDFJSgSKLDFJSKLDFJ" -> False
"ACSKLDFJSGSKLDFJSKLDFJ" -> True
In this Kata, a string is said to be in ALL CAPS whenever it does not contain any lowercase letter so any string containing no letters at all is trivially considered to be in ALL CAPS.
*/

package main

import (
	"fmt"
	"strings"
)

type MyString string

func (s MyString) IsUpperCase() bool {
	cx := strings.ToUpper(string(s))
	sx := string(s)
	switch {
	case cx == sx:
		return true
	default:
		return false
	}
}

func main() {
	wtf := MyString("hello I AM DONALD")
	//wtf := MyString("ACSKLDFJSGSKLDFJSKLDFJ")

	//cx := strings.ToUpper(x)

	fmt.Println(wtf.IsUpperCase())

}

/*
package kata

import (
  "strings"
)

type MyString string

func (s MyString) IsUpperCase() bool {
  return s == MyString(strings.ToUpper(string(s)))
}
*/
