package main

import (
	"fmt"
	"strings"
)

func main() {

	var password string
outer:
	for {
		fmt.Println("Insert password")
		fmt.Scan(&password)
		if strings.Contains(password, "1234") {
			fmt.Println("Weak password. Try again")
			continue
		} else {
			fmt.Println("Password Accepted")
			break outer
		}
	}

}
