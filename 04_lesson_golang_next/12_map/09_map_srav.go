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

	//Сравнение карт . мапы
	//Сравнение масивов: (масси можно использовать как ключ в мапе)
	if [3]int{1, 2, 3} == [3]int{3, 4, 5} {
		fmt.Println("Equel")
	} else {
		fmt.Println("Not equel")
	}

	if [3]int{4, 4, 4} == [3]int{4, 4, 4} {
		fmt.Println("Equel")
	} else {
		fmt.Println("Not equel")
	}

	//Сравнение слайсов // Слайсы сравнивать между сабой нельзя // можно сравнить его только с nil
	//if []int{1,2,3} == []int{1,3,4}

	//Сравнение мапов

	aMap := map[string]int{
		"a": 1,
	}

	bMap := map[string]int{
		"c": 1,
	}

	//if aMap == bMap {}	//Error invalid operation: aMap == bMap (map can only be compared to nil)

	if aMap == nil {
		fmt.Println("Zero map")

	} else {
		fmt.Println("No Zero map")
	}

	if bMap == nil {
		fmt.Println("Zero map")
	} else {
		fmt.Println("No Zero map")
	}

	// Мапа из за того что тип ссылочный можно сравнивать только с nil

	//Грустное следтвие
	//Если мапа/слайс является состовляющей какой-то структорой - структура автаматически не сравнима
}
