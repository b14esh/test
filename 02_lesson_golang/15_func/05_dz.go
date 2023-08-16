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
}
func main() {
	root := squareRoot(-9.3)
		fmt.Printf("\n %0.3f\n", root)
}
