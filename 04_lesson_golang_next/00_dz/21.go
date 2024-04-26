package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var str string = ""

out:
	for {
		fmt.Fscan(os.Stdin, &str)
		if str[0] == ' ' {
			break out
		}
		if strings.HasPrefix(str, "@") {
			break out
		} else {
			fmt.Println(str)
		}

	}
}
