package main

import "fmt"

func main() {
	//Циклы
	// for intit; condition; post {
	// init - блок иницилизации переменных цикла
	// condition - условие (если верно - то тело цикла выполняется, если нет - то цикл завершается)
	// post - изменение переменной цикла (инкрементное действие, декрементарное дейтвие)
	//}

	for i := 0; i <= 5; i++ {
		fmt.Println("Hello world, i =", i)
	}

	///
	for i := 0; i <= 5; i++ {
		fmt.Println("")
	}
	///

	// Важный момент в качестве init может быть использовано выражение ТОЛЬКО чере :=
	// Переменная живет только в теле цикла

	//break - комавнда прерывает выполнения цикла и передает управление  СЛЕДУЮЩИМ ЗА ЦИКЛОМ

	for i := 0; i <= 15; i++ {
		if i > 12 {
			break
		}
		fmt.Printf("Current %d \n", i)
	}
	fmt.Println("After for break")

	///
	for i := 0; i < 4; i++ {
		fmt.Println("")
	}
	///

	//contine - команда, прерывающая выполнение тела цикла и передающая управление СЛЕДУЮЩЕЙ ИТЕРАЦИИ ЦИКЛА

	for i := 0; i <= 20; i++ {
		if i > 10 && i <= 15 {
			continue
		}
		fmt.Printf("Current value: %d\n", i)
	}
	fmt.Println("After for contine")
}
