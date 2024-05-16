package main

import "fmt"

func main() {
	//Строки можно объединять с помощью +.
	fmt.Println("go" + "lang")

	//Целые числа и числа с плавающей запятой.
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	//Булев тип, как вы и ожидаете, с логическими операторами true или false.

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
