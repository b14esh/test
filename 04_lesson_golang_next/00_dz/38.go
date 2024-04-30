//Профессиональная игра в слова. С мягкостью
//Иван и Петр по очереди вводят слова в программу.
// Начиная со второго введенного слова программа анализирует , совпадает ли первая буква нового слова с последней буквой предыдущего слова.
//Если да, то программа работает дальше (считывает еще слово). Если нет - выводит последнее введенное слово и завершает работу.
//В случае, если последняя буква является мягким знаком, то программа сопоставляет первую букву нового слова с буквой, предшествующей мягкому знаку.

/*
Ввод
лань
нож
жир
рыба
ас
сунь
нос
спасите
хобана

Вывод
хобана



Ввод
клюв
вить
ток
конь
ночь
чайник
киров
дирижабль

Вывод
дирижабль

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
			result2 := input2.Text()
			if result2 == "" {
				break
			}
			if input2.Scan() {
				result2 := input2.Text()
				if result2 == "" {
					break
				}

				fullruneSlice1 := []rune(result1)
				fullruneSlice2 := []rune(result2)
				b := fullruneSlice1[len(fullruneSlice1)-1]

				var razmer int

				if len(fullruneSlice1) > 2 {
					razmer = 2

				} else {
					razmer = 1
				}

				x := fullruneSlice1[len(fullruneSlice1)-razmer]
				//x := fullruneSlice1[len(fullruneSlice1)-2]
				//fmt.Println("razmer: ", razmer)

				c := fullruneSlice2[0]

				sb := string(b)
				sx := string(x)
				sc := string(c)
				//fmt.Println(sb, sx, sc)

				if sb == "ь" && sc != sx {
					fmt.Println(0)
					fmt.Println(result1)
					break outline
				} else {
					if sc != sb && sb != "ь" && sc != sx {
						fmt.Println(1)
						fmt.Println(result2)
						break outline
					}

				}

			}
		}
	}
}

//хрень
