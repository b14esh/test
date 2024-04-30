package main

import "fmt"

func main() {

	//Иницилизация
	newMapper := map[string]int{
		"Alice": 1000,
		"Bob":   1000,
	}

	//Получение элемента

	testPerson := "Alice"

	fmt.Println("Salary of testPerson", newMapper[testPerson])

	//Получение элемента который не представлен в мапец

	testPerson = "Derek"

	fmt.Println("Dalaray of new testpeson", newMapper[testPerson]) // При обращение по существующему ключу новая пара не создается

	fmt.Println(newMapper) // нету ошибки как в c++ // новый элемент создан не будет

	//Проверка вхождения ключа
	emploee := map[string]int{
		"Den":    0,
		"Allice": 0,
		"Bob":    0,
	}

	fmt.Println("Den and value", emploee["Den"])
	fmt.Println("Allice and value", emploee["Allice"])
	fmt.Println("Bob and value", emploee["Bob"])
	fmt.Println("Joe and value", emploee["Joe"]) //???

	fmt.Println()
	//Как убедится что Joe нету

	// Мапа при обращение по ключу возращается два значение
	// value и так называемый ok

	if value, ok := emploee["Den"]; ok {
		fmt.Println("Den and value", value)
	} else {
		fmt.Println("Den does not exists in map")
	}

	if value, ok := emploee["Joe"]; ok {
		fmt.Println("JOE and value", value)
	} else {
		fmt.Println("JOE does not exists in map")
	}
}
