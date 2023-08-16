//Ниже приведена программа для вычисления квадратного корня из числа.
//Но если передать функции squareRoot отрицательное число, она вернет ошибку.
//Внесите одно из указанных изменений и попробуйте откомпилировать программу;
//затем отмените изменение и переходите к следующему.
// Посмотрите, что из этого выйдет!
package main

import (
	"fmt"
	"math"
)

func squareRoot(number float64) (float64, error) {
	if number < 0 {
		return 0, fmt.Errorf("can't get square root of negative number")
	}
	return math.Sqrt(number), nil
	//return math.Sqrt(number) // удалили nil
}
func main() {
	//root, err := squareRoot(-9.3)
	root := squareRoot(-9.3) // удалили err
	//root, err := squareRoot(9.3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("\n %0.3f\n", root)
	}
}
