//employed | vacation
//true     | true     => false
//true     | false    => true
//false    | true     => false
//false    | false    => false

/*
func SetAlarm(employed, vacation bool) bool {
	if employed == true && vacation == false {
	  return !(employed && vacation)
	}
	return false
  }


func SetAlarm(employed, vacation bool) bool {
	return employed && !vacation
  }


  func SetAlarm(e, v bool) bool {
	return e && !v
  }

*/

package main

import "fmt"

func SetAlarm(employed, vacation bool) bool {
	e := employed
	v := vacation
	if e == true && v == false {
		return true
	} else {
		return false
	}
}

func main() {

	fmt.Println(SetAlarm(true, false))
	fmt.Println(7032.327531 - 2178.820313)

}
