package main

import (
	"fmt"
	"strings"
)

func main() {

	var password string

	for {
		fmt.Println("Insert password")
		fmt.Scan(&password)
		if strings.Contains(password, "1234") {
			fmt.Println("Weak password. Try again")
		} else {
			fmt.Println("Password Accepted")
			return
		}
	}

}
