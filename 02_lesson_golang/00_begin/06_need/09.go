package main

import "fmt"

func main() {
	if true {
		fmt.Println("Hello go")
	}
	if false {
		fmt.Println("no hello")
	}
	if !false {
		fmt.Println("Za Zu !false")
	}
	if 100 == 100 {
		fmt.Println("Eholo 100 == 100")
	}
}
