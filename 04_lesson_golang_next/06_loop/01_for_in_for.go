package main

import "fmt"

func main() {
	//Вложенные циклы и лейблы
	for i := 0; i < 10; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println("По идеи выше треугольник")

	//
	for i := 0; i <= 3; i++ {
		fmt.Println()
	}
	//

	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			fmt.Printf("i:%d and j:%d and sum i+j=%d \n", i, j, i+j)
		}

	}

	//
	for i := 0; i <= 3; i++ {
		fmt.Println()
	}
	//

	//Иногда бывает плохо
	for i := 0; i <= 2; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("i:%d and j:%d and sum i+j=%d \n", i, j, i+j)
			if i == j {
				break // Хочу что бы все прекратилось / но не выхоит :)
			}
		}

	}

	//
	for i := 0; i <= 3; i++ {
		fmt.Println()
	}
	//

	// лейблы
	//Иногда бывает плохо решение
outer:
	for i := 0; i <= 2; i++ {
		//inner:
		for j := 1; j <= 3; j++ {
			fmt.Printf("i:%d and j:%d and sum i+j=%d \n", i, j, i+j)
			if i == j {
				break outer // А вот так вот мы получим чего хотели
			}
		}

	}

}
