package main

import (
	"fmt"
	"os"
)

func main() {

	if os.Args != nil && len(os.Args) > 3 {
		argsWithProg := os.Args
		argsWithoutProg := os.Args[1:]
		arg := os.Args[3]
		fmt.Println(argsWithProg)
		fmt.Println(argsWithoutProg)
		fmt.Println(arg)
	} else {
		fmt.Println("Enter more argument!\nMin 3 and more!")
	}

}
