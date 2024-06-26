//Часто, когда нам требуется зарегистрироваться на каком-нибудь сервисе, требуется ввести логин и почту. Насколько вам известно, на логин и адрес электронной налагаются определенные требования.
//Напишите программу, которая проверяет, что все хорошо.
//Логин считается валидным (приемлимым), если он содержит не меньше 10 символов и не содержит символа "@".
//Почта считается валидной - если она содержит символ "@" , а также символ точки ".".
//Ваша программа должна реагировать на введенные пользователем данные следующим образом:

//1. если логин короче 10 символов или содержит "@" - вывести "Некорректный логин" и завершиться.
//2. если логин валиден, а адрес почты не содержит "@" или "." - вывести "Некорректная почта" и завершиться.
//3. если и логин и почта - валидны - вывести "ОК" (кириллица) и завершиться.

//Ввод	                   Вывод
//BestPractice_School        ОК
//user@bps.ru

//Best_pr@                   Некорректный логин
//user@bps.ru

package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {

	var (
		login string
		email string
		//login string = "BestPractice_School"
		//email string = "user@bps.ru"
		//login string = "Best_pr@"
		//email string = "user@bps.ru"
	)

	fmt.Println("Пиши: \nBestPractice_School \nuser@bps.ru \n или \nBest_pr@ \nuser@bps.ru")
	fmt.Scan(&login, &email)

	//fmt.Println("Введите логин:")
	//fmt.Scan(&login)
	if utf8.RuneCountInString(login) < 10 || strings.Contains(login, "@") {
		fmt.Println("Некорректный логин")
		return
	}

	//fmt.Println("Введите почту:")
	//fmt.Scan(&email)
	if strings.Contains(email, "@") != true && strings.Contains(email, ".") != true {
		fmt.Println("Некорректная почта")
		return
	}

	fmt.Println("OK")
}
