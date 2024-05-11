package main

import (
	"fmt"
	"strings"
)

func main() {
	s := []string{"I", "am", "a", "special", "one"}
	fmt.Println(strings.Join(s, " ")) // I am a special one
}
