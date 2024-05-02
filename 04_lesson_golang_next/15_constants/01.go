package main

import "fmt"

var PORT = "8000"

//1. Константы это не изменяемые переменные
//1) которые служат для более строго понимания кода
//2) Для того что бы случайно не поменять значения (Предпологается что згачения константы не изменно)
//3) Для удобных преобразований

func main() {
	//2. Обьявление одной константы
	const a = 10
	fmt.Println(a)

	const (
		ipAdress = "127.127.0.3"
		port     = "8000"
		dbname   = "postgress"
	)

	fmt.Println("ipadress value", ipAdress)
}

func checkIpAddress() bool {
	// бужет ошибка
	//.\01.go:27:5: undefined: ipAddress
	// Так как константа бдуте видна только в области видемовсти функции main
	if ipAddress == "127.127.0.3" {
		return true
	}
	return false
}
