// tocelsius преобразует температуру...
package main

import (
	"fmt"
	"keyboard"
	"log"
)

func main() {
	fmt.Print("Enter a F: ")
	fah, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	celsius := (fah - 32) * 5 / 9
	fmt.Printf("%0.2f degrees Celsius\n", celsius)
}
