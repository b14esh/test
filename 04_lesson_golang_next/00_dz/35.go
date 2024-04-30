//Знако-переменный ряд

//Напишите программу, которая считывает число n ∈ N, затем n чисел (целых), каждое с новой строки.

//После чего программа вычисляет и печатает знакочередующуюся сумму ряда (первое число прибавить, второе вычесть, третье прибавить, четвертое вычесть, и т.д.) .
//Например, для чисел 1,2,3,4,5  сумма будет выглядеть следующим образом:
//1 - 2 + 3 - 4 + 5 = 3 // ОК сделал
//Ввод	Вывод ????
//3		5
//5
//5
//5

// Пример 2
// Ввод	Вывод?????
// 4		-2
// 1
// 2
// 3
// 4

package main

import "fmt"

func main() {

	var i int
	var x int = 0
	var minus int = -1
	var flag int = 0
	var smen int
	var vsum int
	box := []int{}

	for {
		fmt.Scan(&i)
		if flag == 0 {
			box = append(box, i)
			x++
			flag = 1
		} else if flag == 1 {
			smen = (i * minus)
			box = append(box, smen)
			x++
			flag = 0

		}
		if x == 5 {
			for _, val := range box {
				vsum += val
			}
			fmt.Println(box)
			fmt.Println("Сумма веденного борохла равна:", vsum)
			break

		}

	}

	//???? хз
}
