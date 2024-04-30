//Знако-переменный ряд

//Напишите программу, которая считывает число n ∈ N, затем n чисел (целых), каждое с новой строки.

//После чего программа вычисляет и печатает знакочередующуюся сумму ряда (первое число прибавить, второе вычесть, третье прибавить, четвертое вычесть, и т.д.) .
//Например, для чисел 1,2,3,4,5  сумма будет выглядеть следующим образом:
//1 - 2 + 3 - 4 + 5 = 3
//Ввод	Вывод
//3		5
//5
//5
//5

//Пример 2
//Ввод	Вывод
//4		-2
//1
//2
//3
//4

package main

import "fmt"

func main() {

	//var num int
	//var x int
	//var a int
	//var b int
	//var c int
	//var d int
	//var y int

	//bookslice := []int{}

	//x = 1 - 2 + 3 - 4 + 5 //OK 3
	//x = 3 - 5 + 5 - 5  //??? 5?????
	//x = 5 - 5 + 5 - 5 + 5
	//x = 4 - 1 + 2 - 3 - 4 //OK
	//fmt.Println(x)

	// Прикольное поведение :)
	//outline:
	//	for {
	//		fmt.Scanln(&num)
	//		if num == 0 {
	//			break outline
	//		}
	//		if num > 0 {
	//			if num == 0 {
	//				break outline
	//			}
	//			bookslice = append(bookslice, num)
	//		}
	//		for _, vol := range bookslice {
	//			if num == 0 {
	//				break outline
	//			}
	//			fmt.Printf("%d ", vol)
	//		}
	//
	//	}

	/* Не решение но эмитация
	   outline:
	   	for {
	   		fmt.Scanln(&num)
	   		if num > 0 {
	   			//fmt.Println(num)
	   			bookslice = append(bookslice, num)
	   			num = 0
	   			x++
	   		}
	   		if x == 2 {
	   			a = bookslice[0] - bookslice[1]
	   			//fmt.Println(a)
	   			//break outline

	   		}
	   		if x == 3 {
	   			b = a + bookslice[2]

	   		}
	   		if x == 4 {
	   			c = b - bookslice[3]

	   		}
	   		if x == 5 {
	   			d = c + bookslice[4]
	   			fmt.Println(d)
	   			break outline
	   		}
	   	}
	*/

	// Забавно что то порешал но это все не то
	var a int
	var b int
	var x int
	var flag int
	var z int
	box1 := []int{}
	box2 := []int{}

	for {
		fmt.Scan(&a, &b)
		box1 = append(box1, a)
		box2 = append(box2, b)
		x++
		if x > 1 {
			break
		}

	}

	for _, val := range box1 {
		fmt.Printf("%d ", val)
	}
	fmt.Println()
	for _, val := range box2 {
		fmt.Printf("%d ", val)
	}
	fmt.Println()
	for _, val1 := range box1 {
		for _, val2 := range box2 {
			if flag == 0 {
				z = val1 + val2
				flag = 1
				fmt.Println("Плюс ", z)
			}
			if flag == 1 {
				z = val1 - val2
				flag = 0
				fmt.Println("Минус", z)

			}
		}
	}

	// пока не порешал :(
}
