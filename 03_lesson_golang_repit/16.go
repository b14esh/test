package main 

import "fmt"

func main() {

	// Указатели не используем
	var x = 0
	pointer (x)
	fmt.Println(x)

}


func pointer (x int) {
	x = 2 
}
