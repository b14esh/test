package main

import "fmt"

func main() {
	var nilMap map[int]string
	fmt.Printf("%#v\n", nilMap)
	nilMap[3] = "three"
}
