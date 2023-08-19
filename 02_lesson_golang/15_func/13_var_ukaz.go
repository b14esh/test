package main

import "fmt"

func main() {
	var myBool bool
	myBoolPointer := &myBool //Короткое объявление переменной-указателя.
	fmt.Println(myBoolPointer)
}
