package main

import (
	"fmt"
	"strings"
)

func main() {

	//Млдификация цикла for
	//1. Классический цикл "wile do"

	//var loopVar int = 0
	// примерно так выглядит в других яхыках
	//while (loopVar = 10)
	//	...loopvar
	//	loopVar++

	var loopVar int = 0
	for loopVar < 10 {
		fmt.Println("In while like loop %d", loopVar)
		loopVar++

	}

	//2. Классический бесконечный цикл
	//for {
	//fmt.Println("KEK")
	//}
