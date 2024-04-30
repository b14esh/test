package main

import "fmt"

func main() {

	// Игнорировани index в range for цикле
	floatArr := [...]float64{10.1, 15.4, 44.44, 12.5}
	// _ - специальная переменная на которую GO не будет ругатся что она не проиницилизирована

	for index, _ := range floatArr {
		fmt.Printf("%d value index \n", index)
	}
}
