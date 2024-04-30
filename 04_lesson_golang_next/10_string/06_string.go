package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	//Длина и емкасть строки
	// Длина len() - количество байт в слайсе

	fmt.Println("Lenght of Вася", len("Вася"), "bates") // Но тут есть облом ;)

	// Длина RuneCounter - количесто !РУН!

	fmt.Println("Lenght of Вася", utf8.RuneCountInString("Вася"), "runes") // вот так мы узнаем точную длину !!!  и только так!!

	// Вычисление емкасти строки бессмысленно, т.к. строка базовый тип
	// не будет работать fmt.Println(cap("Вася"))
	// можно исхетрится
	//fmt.Println(cap([]rune("Вася"))) //Но это не нужно делать

}
