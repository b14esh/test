//Напишите программу, которая проверяет правильность расчета на "раз два три".

//Пользователь вводит последовательно 3 строки.

//если эти строки "раз", "два", "три" - вывести "ОК" (кириллица)
//если вместо строки "раз" введена "один" - вывести "ОК" (кириллица)
//если вместо всех слов введены соответствующие числа "1", "2", "3" - вывести "ОК" (кириллица)
//"НЕ ПРАВИЛЬНО" - во всех остальных случаях

//раз            ОК
//два
//три

//один           ОК
//два
//три

package main

import (
	"fmt"
	"strings"
)

func main() {

	var (
		one   string
		two   string
		three string
	)

	fmt.Println("Введите: \nПример1: \n\t раз два три \nПример2: \n\t один два три \nПример3: \n\t 1 2 3")
	fmt.Scan(&one, &two, &three)

	if strings.Contains(one, "один") || strings.Contains(one, "раз") || strings.Contains(one, "1") {
		if strings.Contains(two, "два") || strings.Contains(two, "2") {
			if strings.Contains(three, "три") || strings.Contains(three, "3") {
				fmt.Println("ОК")
			}
		}

	} else {
		fmt.Println("НЕ ПРАВИЛЬНО")
	}

}
