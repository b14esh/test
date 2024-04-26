//Простой калькулятор

package main

import "fmt"

func printerros() {
	fmt.Printf("\nОшибка! Ввод не верный!  \nПример: \n\t10 - 10\n\t10 + 10\n\t10 * 10\n\t10 \\ 10")

}

func main() {

	var a float64
	var b float64
	var c string
	var x float64

	fmt.Println("Пример: 10 + 10")
	//fmt.Scan(&a, &c, &b)
	fmt.Scanln(&a, &c, &b)
	switch {
	case c == "/":
		if b == 0 {
			fmt.Println("На ноль делить не льзя")
			return

		}
		x = a / b
		fmt.Printf("%0.2f / %0.2f = %0.2f \n", a, b, x)
	case c == "*":
		x = a * b
		fmt.Printf("%0.2f * %0.2f = %0.2f \n", a, b, x)
	case c == "-":
		x = a - b
		fmt.Printf("%0.2f - %0.2f = %0.2f \n", a, b, x)
	case c == "+":
		x = a + b
		fmt.Printf("%0.2f + %0.2f = %0.2f \n", a, b, x)
	case c == " ":
		printerros()

	default:
		printerros()
	}

}
