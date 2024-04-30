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

	// У мапы есть длина, кол-во пар это и есть длина мапы

	fmt.Println("Pair amount in map", len(emploee))

	// Мапа как и салйс ссылочный тип

	words := map[string]string{
		"One": "Один",
		"Two": "Два",
	}

	newwords := words
	newwords["three"] = "Три"
	delete(newwords, "One")
	fmt.Println("wordpam", words)
	fmt.Println("newwords", newwords)

}
