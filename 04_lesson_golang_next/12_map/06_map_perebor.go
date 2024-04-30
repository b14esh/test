package main

import "fmt"

func main() {
	//Перебор карты / мапы
	// Перебор элементов карты / Этерирование
	//Проверка вхождения ключа

	// ЕСЛИ НАДА ПУСТАЯ КАРТА ИСПОЛЬЗУЙТЕ make
	//	mapp := make(map[string]int)

	emploee := map[string]int{
		"Den":    56,
		"Allice": 11,
		"Bob":    222,
	}

	for key, value := range emploee {
		fmt.Printf("%s and value %d\n", key, value)
	}

}
