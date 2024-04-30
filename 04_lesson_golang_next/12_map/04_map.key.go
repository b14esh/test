package main

import "fmt"

func main() {

	// Map , карта, это паразначений ключ и значение. Иницилизация пустой карты / мапы
	mapper := make(map[string]int)

	fmt.Println("Empty map:", mapper)

	//Добавление пар в существующую мапу
	mapper["Alice"] = 24
	mapper["Bob"] = 25
	fmt.Println("Mappert afte add", mapper)

	//Иницилизация
	newMapper := map[string]int{
		"Alice": 1000,
		"Bob":   1000,
	}

	newMapper["Jo"] = 3000
	fmt.Println("New Mapper:", newMapper)

	//Что может быть ключем в мапе
	//ВАЖНО: МАПА НЕ УПОРЯДОЧЕНА В GO
	//КЛЮЧЕМ МОЖЕТ БЫТЬ ТОЛЬКО СРАВНИМЫЙ ТИП (==, !=)
	//НУЛЕВОЕ ЗНАЧЕНИЕ ДЛЯ MAP
	//var ZeroValie map[string]int // Zerovalie == nil
	//ZeroValie["Alice"] = 12      //panic: assignment to entry in nil map
	//Используйте команду make для создания мапы

	//

}
