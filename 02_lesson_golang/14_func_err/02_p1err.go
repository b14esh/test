package main

import (
	"errors"
	"fmt"
)

func errorsi() {
	err := errors.New("error lol")
	fmt.Println(err.Error())
        fmt.Printf("\nHi %s, LOL %d Gramm, hor %3.3f zz,  ", "hohoh", 100, 5.56)
        fmt.Printf("\nHi %s, LOL %d Gramm, hor %1.1f zz,  ", "hohoh", 100, 5.56)
}

func main() {

	x1 := 10

	for x := 1; x <= 10; x++ {
		errorsi()
		fmt.Println(x1 - x)
	}
}
