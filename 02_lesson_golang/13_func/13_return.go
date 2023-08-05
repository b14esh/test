package main

import "fmt"

func double(number float64) float64 {
	return number * 2
}
func main() {
	dozen := double(6.0) //Возвращаемое значение присваивается переменной
	fmt.Println(dozen)
	fmt.Println(double(4.2)) //Возвращаемое значение передается другой функции.
}
