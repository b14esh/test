//average вычисляет среднее значение
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// GetFloats читает значение float64 из каждой строки файла. ошибку.
func GetFloats(fileName string) ([3]float64, error) {
	var numbers [3]float64 // Объявление возвращаемого массива.

	file, err := os.Open(fileName) //Открывает файл с переданным именем.
	if err != nil {
		return numbers, err
	}

	i := 0 //Переменная для хранения индекса, по которому должно выполняться присваивание

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64)

		if err != nil {
			return numbers, err
		}
		i++ //Переход к следующему индексу массива.
	}

	err = file.Close()
	if err != nil {
		return numbers, err
	}

	if scanner.Err() != nil {
		return numbers, scanner.Err()
	}
	return numbers, nil //Если выполнение дошло до этой точки, значит, ошибок не было, поэтому программа возвращает массив чисел и значение ошибки «nil».
}

func main() {
	filename := "data1.txt"
	//filename := "data.txt"
	numbers, err := GetFloats(filename)
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	sampleCount := float64(len(numbers))
	fmt.Printf("Average: %0.2f\n", sum/sampleCount)
}
