package main

import "fmt"

func main() {
	var myStruct struct { // Объявление переменной с именем "myStruct".
		//Переменная «myStruct» может хранить структуры, состоящие из поля float64 с именем «number», поля string с именем «word» и поля bool с именем «toggle».
		number float64
		world  string
		toggle bool
	}
	fmt.Printf("%#v\n", myStruct)
}
