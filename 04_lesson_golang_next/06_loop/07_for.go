package main

import "fmt"

func main() {

	//Самый базовый тип, с единственным условием.
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	//Классический цикл оператора for начальное значение/условие/постусловие.
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	//for без условия будет работать до тех пор, пока вы не примените break или return внутри функции.
	for {
		fmt.Println("loop")
		break
	}
}
