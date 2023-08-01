package main

import "fmt"

func main() {
	for x := 1; x <= 3; x++ {
		fmt.Println("before continue")
		continue
		fmt.Println("after continue") //эту строку мы никогда не увидим
	}

}
