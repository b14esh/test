/*Ниже приведен код функции negate, которая должна обновлять значение переменной truth и заменять его противоположным (false),
а также обновлять значение переменной lies и заменять его противоположным (true).
Но если вызвать negate для переменных truth и lies и вывести их значения, становится ясно, что они не изменились!
*/

package main

import "fmt"

func negate(myBoolean bool) bool {
	return !myBoolean
}

func main() {
	truth := true
	negate(truth)
	fmt.Println(truth)
	lies := false
	negate(lies)
	fmt.Println(lies)
}
