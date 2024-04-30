package main

import "fmt"

func main() {

	//Создание строки из слайса байт
	myBytesSlice := []byte{0x40, 0x41, 0x42, 0x43}

	mystr := string(myBytesSlice)

	fmt.Println(mystr)

	// в десятичном
	myDecByt := []byte{100, 101, 102, 103}
	myDecStr := string(myDecByt)
	fmt.Println(myDecStr)

	// создание строки из рун
}
