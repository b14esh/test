// Пакет datafile предназначен для чтения данных из файлов.
//package datafile
package main

import (
	"fmt"
	"bufio"
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
	fileName := "data.txt"
	fmt.Println("Hello readfile")
	fmt.Println(GetFloats(fileName))
}
