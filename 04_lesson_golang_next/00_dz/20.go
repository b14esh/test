package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	//str := ""
	//fmt.Println(unsafe.Sizeof(str))

	//emptystring := ""
	//if len(emptystring) == 0 {
	//fmt.Println("String in Empty")
	//}

out:
	for {
		str := ""
		//fmt.Scan(&str)
		fmt.Fscan(os.Stdin, &str)
		//fmt.Println(unsafe.Sizeof(str))
		switch {
		case strings.Contains(str, "1234"):
			break out
		case len(str) == 0:
			break out

		//case unsafe.Sizeof(str) == 16:
		//	break out

		case len(str) > 1:
			fmt.Println(str)
		default:
			break out

		}

	}

}
