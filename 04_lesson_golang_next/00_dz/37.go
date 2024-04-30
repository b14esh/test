//Профессиональная игра в слова
//Иван и Петр по очереди вводят слова в программу.
//Начиная со второго введенного слова программа анализирует , совпадает ли первая буква нового слова с последней буквой предыдущего слова.
//Если да, то программа работает дальше (считывает еще слово). Если нет - выводит последнее введенное слово и завершает работу.

//Ввод:
//мяч
//чай
//йод
//дом
//муха
//орел
//Вывод
//орел
/*
Ввод
орел
лук
картошка
ананас
собрание
трус
Вывод
трус
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	input1 := bufio.NewScanner(os.Stdin)
	input2 := bufio.NewScanner(os.Stdin)
outline:
	for {
		if input1.Scan() {
			result1 := input1.Text()
			if result1 == "" {
				break
			}
			if input2.Scan() {
				result2 := input2.Text()
				if result2 == "" {
					break
				}

				fullruneSlice1 := []rune(result1) // преобразовываем  result в руны
				fullruneSlice2 := []rune(result2) // преобразовываем  result в руны

				//fmt.Printf("%c \n", fullruneSlice1[0])
				//fmt.Printf("%c \n", fullruneSlice1[len(fullruneSlice1)-1])
				//fmt.Printf("%c \n", fullruneSlice2[0])
				//fmt.Printf("%c \n ", fullruneSlice2[len(fullruneSlice2)-1])

				// немного упращаем дальнейший перебор в if
				//a := fullruneSlice1[0]                     // первая буква первого слова
				b := fullruneSlice1[len(fullruneSlice1)-1] // последняя буква первого слова
				c := fullruneSlice2[0]                     // первая буква второго слова
				//d := fullruneSlice2[len(fullruneSlice2)-1] // последняя буква второго слова

				if b != c {
					fmt.Println(result2)
					break outline

				}

			}
		}
	}
}
