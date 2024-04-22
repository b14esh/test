package main 

import "fmt"

func main() {

	// Указатели используем и получаем наше число 2 из функции pointer
	var x = 0
	pointer (&x)
	fmt.Println(x)

}


func pointer (x *int) {
	*x = 2 
}
