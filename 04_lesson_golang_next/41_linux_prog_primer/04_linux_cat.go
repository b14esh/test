// Влинукс
// cat -n test.txt
// Перед нами — одна из самых простых Bash-команд.
//Она просто читает файл и выводит на экран его содержимое. Опция -n обеспечивает вывод перед строками их номеров.

package main

import (
	"fmt"
	"os"
)

/*
Основной объём вышеприведённого кода используется для вывода номеров строк,
так как для вывода исходного содержимого файла нам достаточно прибегнуть к fmt.Println.
Мы используем функцию fmt.Sprintf для создания префиксов в начале строк.
Префикс содержит число, выровненное по правому краю,
то есть — вывод получится таким же, как при использовании команды cat.
*/

func cat(fileName string, printLineNumbers bool) {
	fileData, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	if printLineNumbers {
		outputData := make([]byte, 0, len(fileData))

		var lineCount int

		for i, d := range fileData {
			if d == '\n' || i == 0 {
				if i != 0 {
					outputData = append(outputData, d)
				}

				lineCount++

				startOfLineString := fmt.Sprintf("%6d  ", lineCount)

				outputData = append(outputData, startOfLineString...)

				if i == 0 {
					outputData = append(outputData, d)
				}
			} else {
				outputData = append(outputData, d)
			}
		}

		fmt.Printf("%s\n", outputData)
	} else {
		fmt.Println(string(fileData))
	}
}

func main() {
	cat("test.txt", true)
}
