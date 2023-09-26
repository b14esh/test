package main

import "fmt"

func probels(i int) {
	if i == 0 {
		 fmt.Printf("\t")
	}
	if i == 1 {
		fmt.Printf("\t\t")
	}
	if i == 2 {
		fmt.Printf("\t\t\t")
	}
	if i == 3 {
		fmt.Printf("\t\t\t\t")
	}
	if i == 4 {
		fmt.Printf("\t\t\t\t\t")
	}
	return
}
func main() {
	var xxx [8]string
	xxx[0] = "Hel"
	xxx[1] = "lo"
	xxx[2] = " ^_^ "
	xxx[3] = "GO"
	xxx[4] = "!!!"
	xxx[5] = "END"

	for i := 0; i < 5; i++ {
		y := i
		probels(y)
		//fmt.Println(xxx[i])
		fmt.Printf(xxx[i])
	}
}
