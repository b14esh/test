package main

import "fmt"

func main() {

	// Игнорировани id в range for цикле
	floatArr := [...]float64{10.1, 15.4, 44.44, 12.5}
	// _ - специальная переменная на которую GO не будет ругатся что она не проиницилизирована

	for _, val := range floatArr {
		fmt.Printf("%.2f value WO id\n", val)
	}
}
