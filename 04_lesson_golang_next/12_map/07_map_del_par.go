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

	//Как удалить пары
	//Удаление существующей пары
	fmt.Println("Befor del", emploee)
	delete(emploee, "Den")
	fmt.Println("After deleting", emploee)

	//Что будет если удалить не существующего
	fmt.Println("Befor1 del", emploee)
	delete(emploee, "Anna")
	fmt.Println("After1 deleting", emploee)

	// Если проверка не была сделана то ничего не удаляется
	if _, ok := emploee["Anna"]; ok { //Проверь
		delete(emploee, "Anna") //ОЕНЬ РЕСУРСОЕМКАЯ ОПЕРАЦИЯ
	} else {
		fmt.Println("Anna does not exists in map")
	}
}
