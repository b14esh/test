package main

import "fmt"

func main() {

	// создание строки из рун
	// руны как hex

	runeHexSlice := []rune{0x45, 0x46, 0x47, 0x48}
	mystrFromeRune := string(runeHexSlice)

	fmt.Println("From Rune(hex):", mystrFromeRune)

	//руны как литералы (сразу как символы)

	runeliteralSlice := []rune{'V', 'a', 's', 'y', 'a'} // ' ' таким сиволом обозначается руна
	mystrFromeRuneLiteral := string(runeliteralSlice)
	fmt.Printf("%s from rune kiterals %T", mystrFromeRuneLiteral, mystrFromeRuneLiteral)

}
