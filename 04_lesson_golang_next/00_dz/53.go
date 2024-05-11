/*
Write a function to split a string and convert it into an array of words.

Examples (Input ==> Output):
"Robin Singh" ==> ["Robin", "Singh"]

"I love arrays they are my favorite" ==> ["I", "love", "arrays", "they", "are", "my", "favorite"]
*/

package main

import (
	"fmt"
	"strings"
)

func StringToArray(str string) []string {

	words := strings.Fields(str)
	box := []string{}
	sbox := []string{}

	for i := 0; i < len(words); i++ {
		box = append(box, words[i])
	}

	sbox = box

	return sbox

}

func main() {
	i := "I love arrays they are my favorite"
	//i = "Robin Singh"
	fmt.Println(StringToArray(i))

}

/*
package kata

import "strings"


func StringToArray(str string) []string {
  words := strings.Fields(str)
	box := []string{}
	sbox := []string{}

	for i := 0; i < len(words); i++ {
		box = append(box, words[i])
	}

	sbox = box

	return sbox
}
*/

/*
package kata

import "strings"

func StringToArray(str string) []string {
      return strings.Fields(str)
}
*/

/*
package kata

import "strings"

func StringToArray(str string) []string {
      a := strings.Split(str, " ")
      return a
}

*/

/*
package kata
import "strings"


func StringToArray(str string) []string {


 word := strings.Fields(str)


  return word
      // your code here
}

*/

/*
package kata

import "strings"

func StringToArray(str string) (result []string) {
  result = strings.Fields(str)

  return
}

*/

/*
package kata

import (
  "bufio"
  "strings"
)

func StringToArray(str string) []string {
      arr := []string{}
      strReader := strings.NewReader(str)
      scanner := bufio.NewScanner(strReader)
      scanner.Split(bufio.ScanWords)
      for scanner.Scan() {
        arr = append(arr, scanner.Text())
      }
      return arr
}

*/
