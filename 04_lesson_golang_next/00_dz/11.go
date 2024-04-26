package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("Hello world!")

	//Анализатор работает следующим образом :
	//если пользователь вводит произвольную строку и в данной строке встречается подстрока "чат" требуется вывести "БОТ",
	//и "НЕ БОТ" в противном случае.

	//Пример:
	//Початимся? БОТ

	//Поиск подстроки в строке
	//totalString, substring := "ABCDEDFG", "CDE" //true
	//totalString, substring := "ABCDEDFG", "cde" //false
	//fmt.Println(strings.Contains(где_ищем, что_ищем))
	//fmt.Println(strings.Contains(totalString, substring))

	//strings.Compare
	//fmt.Println(strings.Compare("a", "abcd"))
	//0 если  a == b
	//-1 если a < b
	//1 если  a > b

	var chat string

	fmt.Println("Напишите: \t\t\tПочатимся?")
	fmt.Scan(&chat)

	//if strings.Compare(chat, "чат") > 0 || strings.Compare(chat, "чат") == 0 || strings.Contains(chat, "чат") == true && strings.Compare("чат", chat) < 0 {
	//  fmt.Println("БОТ")
	//} else {
	//	fmt.Println("НЕ БОТ")
	//}

	if strings.Contains(chat, "чат") == true {
		fmt.Println("БОТ")
	} else {
		fmt.Println("НЕ БОТ")
	}
}
