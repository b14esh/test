// average2 вычисляет среднее значение.
// use: go run 01_average2.go 50 10 55 66 99 45 22 33 100
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args[1:]
	var sum float64 = 0
	for _, argument := range arguments {
		number, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			log.Fatal(err)
		}
		sum += number //Число прибавляется к сумме
	}
	sampleCount := float64(len(arguments))
	fmt.Printf("Average: %0.2f\n", sum/sampleCount) //Вычисление среднего значения
}
