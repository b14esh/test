//Высокая математика
//Напишите программу, которая считывает одно число n ∈ N - после чего, для каждого из целых чисел от -n до n (включительно) выводится фраза: "Квадрат числа [число] равен [число в квадрате]".
//Фраза выводится для каждого числа с новой строки.

//Ввод
//2

//Вывод
//Квадрат числа -2 равен 4
//Квадрат числа -1 равен 1
//Квадрат числа 0 равен 0
//Квадрат числа 1 равен 1
//Квадрат числа 2 равен 4

package main

import "fmt"

func main() {
	var (
		a int
		c int
		x int
		y int
	)
	plusi := []int{}
	minusi := []int{}
	supermass := []int{}

	fmt.Scan(&a)
	//a = 4
	c = a

	//for i := a; i >= 0; i-- {
	//	fmt.Println(a)
	//	plusi = append(plusi, a)
	//	a--

	//}

	for i := 0; i <= int(a); i++ {
		//fmt.Println(i)
		plusi = append(plusi, i)

	}

	x = -1 * c

	for i := x; i < 0; i++ {
		//fmt.Println(i)
		minusi = append(minusi, x)
		x++

	}

	supermass = append(supermass, minusi...)
	supermass = append(supermass, plusi...)

	//fmt.Println(supermass)

	for _, val := range supermass {
		y = val * val
		fmt.Printf("Квадрат числа %d равен %d \n", val, y)
	}

}
