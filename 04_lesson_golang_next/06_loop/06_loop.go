package main

import "fmt"

func main() {
	//Цикл с множественными переменными цикла

	// не будет работать
	//for x, y := 0, 1; x <= 10 && y <= 12; x++ ; y++ { // x++ и y ++ мы получим не то что нам нужно / и получается по итогу x++, y++ ->> x++; y++, так как будет не верная запись и y++ отсекется
	//	fmt.Printf("%d + %d = %d\n", x, y, x+y)
	//}

	// рабочий вариант
	for x, y := 0, 1; x <= 10 && y <= 12; x, y = x+1, y+2 {
		fmt.Printf("%d + %d = %d\n", x, y, x+y)

	}
}
