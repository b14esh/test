package main

import "fmt"

func double(number float64) float64 {
	return number * 2       //Функция всегда должна завершаться здесь...
	fmt.Println(number * 2) //А эта строка никогда не должна выполняться!
}
