package main

import "fmt"

func main() {

	var i int

	// Гадость с терминацией for из switch
uberloop:
	for {
		fmt.Scan(&i)
		switch {
		case i%2 == 0:
			fmt.Printf("Value is %d and it's even\n", i)
			break uberloop
		}
	}
	fmt.Println("End")

}

// СТРАЙТЕСЬ ВСЕГДА ИСПОЛЬЗОВАТЬ ЛАЙБЛЫ кода у вас имеестя много вложений с циклом for
