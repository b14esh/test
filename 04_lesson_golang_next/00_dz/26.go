//Делим десятичное число на 2 и записываем остаток от деления.
//Результат деления вновь делим на 2 и опять записываем остаток.
//Повторяем операцию до тех пор пока результат деления не будет равен нулю.
//Полученные остатки будут представлять двоичное число.

package main

import (
	"fmt"
)

func main() {
	var (
		a int64
		b int64
		c int64
		x int64
	)

	slicex := []int64{}

	//a = 4
	//a = 10 //1010
	//a = 12 //1100
	//a = 13 //1101
	//fmt.Println(strconv.FormatInt(a, 2))

	fmt.Scan(&a)
	fmt.Println("Вы ввели:", a)

	c = a % 2
	slicex = append(slicex, c)

outline:
	for {
		if a != 0 {
			b = a / 2
			a = b
			c = b % 2
			fmt.Printf("Текщие значение a = %d, проверка остатка = %d \n", a, c)
			if a != 0 {
				slicex = append(slicex, c)
			}
			x++
		}
		if a == 0 {
			fmt.Printf("Делили на два = %d раз/a.\n", x)
			fmt.Println("Получили массив :", slicex)

			fmt.Println("Читаем массив в строку:")
			for _, val := range slicex {
				fmt.Printf("%d", val)
			}

			//Вот так вот можно читать массив
			//for i := 0; i < len(floatArr); i++ {
			//	fmt.Println(floatArr[i])
			//}

			fmt.Println("\nВот этого мы хотели добится:")
			for i := len(slicex) - 1; i >= 0; i-- {
				fmt.Printf("%d", slicex[i])
			}

			break outline
		}

	}
}
