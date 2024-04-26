package main

import "fmt"

func main() {

	//Кейсы с проваливаниями
	var number int
	fmt.Scan(&number)

	switch {
	case number < 100:
		fmt.Printf("%d is less then 100\n", number)
		fallthrough // проволится дальше // Проваливание выполняют ДАЖЕ ЛОЖНЫЕ КЕЙСЫ // Не проверяются условия следущего кейса
	case number < 200:
		fmt.Printf("%d is less then 200\n", number)
		fallthrough // проволится дальше // Проваливание выполняют ДАЖЕ ЛОЖНЫЕ КЕЙСЫ // Не проверяются условия следущего кейса
	case number > 200:
		fmt.Printf("%d is GREATER then 200\n", number)
		fallthrough // проволится дальше // Проваливание выполняют ДАЖЕ ЛОЖНЫЕ КЕЙСЫ // Не проверяются условия следущего кейса
	default:
		fmt.Println("OOO", number)

	}

}
