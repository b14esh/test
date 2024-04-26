package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {

	//strings
	//Строка это последовательность чего-то, заключенная в кавычки

	name := "Vasia"
	rusname := "Вася"
	lastname := "Pupkin"
	concat := name + " " + lastname
	fmt.Println("Full name:", concat)

	// Дина строки
	fmt.Println("lenght of string:", name, len(name))
	// Длина строки для русских букв
	fmt.Println("lenght of string:", rusname, len(rusname))
	//Строка это последовательность чего-то, заключенная в кавычки
	// А вот и нет!
	// Строка в Go это набор байт
	// Функция len возращает кол-во элементов в наборе

	// А вот так вот можно определить кол-во сиволов
	fmt.Println("Amaunt of chars:", utf8.RuneCountInString(rusname))
	//Rune один UTF символ

	//Поиск подстроки в строке
	//totalString, substring := "ABCDEDFG", "CDE" //true
	totalString, substring := "ABCDEDFG", "cde" //false
	//fmt.Println(strings.Contains(где_ищем, что_ищем))
	fmt.Println(strings.Contains(totalString, substring))

	//rune -> alias int32
	var sampleRune rune
	var anotheRune rune = 'Q' // Одинарные ковычки важны
	var thirdRune rune = 1234567
	var agRune = 88

	//какому сиволу соответствует rune
	fmt.Printf("Rune as char %c \n", sampleRune)
	fmt.Printf("Rune as char %c \n", anotheRune)
	fmt.Printf("Rune as char %c \n", thirdRune)
	fmt.Printf("Rune as char %c \n", agRune)

	//"A" < "abcd" / говорят что в го пока плохо работает и лучше использовать функцию string.Compare
	fmt.Println("A" < "abcd")
	fmt.Println("a" < "abcd")

	//strings.Compare
	fmt.Println(strings.Compare("a", "abcd"))
	//0 если  a == b
	//-1 если a < b
	//1 если  a > b

	var aByte byte // alias int 8
	fmt.Println("abyte:", aByte)
}
