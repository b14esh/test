package main

import "fmt"

// Свои функции пример использования

func main() {
 var a = 29
 var b = 1
 var r int
 r = summ(a,b)
 fmt.Println(r)

}

//   summ  имя функции
//   num_1 первый параметр передаваемый функции
//   num_2 второй параметр передоваемый функции
//   функция возрощает int (целое число)

func summ(num_1 int, num_2 int) int {
	var res int
	res = num_1 + num_2
	return res
}
