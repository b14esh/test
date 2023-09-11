//Отложенные вызовы функций помещаются в стек.
//Когда функция завершает работу, ее отложенные вызовы выполняются в порядке «последним пришел — первым вышел».

package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
