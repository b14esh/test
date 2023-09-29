// average2 вычисляет среднее значение.
// use: go run 09_average2_many_strings.go 23.3 34.7 77.4  99
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func average(numbers ...float64) float64 {
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
return sum / float64(len(numbers))
}

func main() {
	arguments := os.Args[1:]
	//var sum float64 = 0
	var numbers []float64
	for _, argument := range arguments {
		number, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			log.Fatal(err)
		}
		//sum += number //Число прибавляется к сумме
		numbers = append(numbers, number)
	}
	//sampleCount := float64(len(arguments))
	//fmt.Printf("Average: %0.2f\n", sum/sampleCount) //Вычисление среднего значения
	fmt.Printf("Average: %0.2f\n", average(numbers...)) //Вычисление среднего значения
}
