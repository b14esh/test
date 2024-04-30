//Начинающий исследователь

//Вы несколько раз подбрасываете монету и вас интересует, какое максимальное количество орлов может выпасть.
//Напишите программу, которая принимает на вход строку, состоящую из букв "р" и "о" (кириллица) - символизирующих выпадение решки и орла соответственно.

//На выходе программа должна ответить на вопрос - каково максимальное число выпавших подряд орлов?

//ороророророоррооооор // 5

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	/*
	   	str := "абвбааааба"

	   	maxChar := ' '
	   	maxCount := 0

	   	currentChar := ' '
	   	currentCount := 0

	   	for _, char := range str {

	   		if char == currentChar {
	   			currentCount++
	   		} else {
	   			currentChar = char
	   			currentCount = 1
	   		}

	   		if currentCount > maxCount {
	   			maxCount = currentCount
	   			maxChar = currentChar
	   		}
	   	}
	   	fmt.Printf("Максимальное скопление символа %c: %d\n", maxChar, maxCount)
	   }

	*/

	input := bufio.NewScanner(os.Stdin)
	for {
		if input.Scan() {
			result := input.Text()

			if result == "" {
				break
			}

			box := []rune(result)
			box_orel := []rune{}
			box_orel_read := []rune{}
			var bo_count int
			box_orel_null := []rune{}
			var max_count_orel int

			for i := 0; i < len(box); i++ {
				box_orel = append(box_orel, box[i])
			}

			//fmt.Println(len(box_orel))

			for i := 0; i < len(box_orel); i++ {
				if box_orel[i] == 'р' && bo_count > 1 {
					max_count_orel = len(box_orel_read)
					box_orel_read = box_orel_null
					bo_count = 0

				}

				if box_orel[i] == 'о' {
					box_orel_read = append(box_orel_read, box_orel[i])
					bo_count++

				}

			}

			//fmt.Println(len(box_orel_read))
			//fmt.Println(bo_count)
			//fmt.Println(max_count_orel)
			if len(box_orel_read) == 0 && bo_count == 0 && max_count_orel == 0 {
				fmt.Println(0)
				break
			}

			if len(box_orel_read) == bo_count && max_count_orel == 0 {
				fmt.Println(bo_count)
			}
			if len(box_orel_read) == 0 && bo_count == 0 {
				fmt.Println(max_count_orel)
			}
			if len(box_orel_read) == 0 && bo_count == 0 && max_count_orel == 0 {
				fmt.Println(0)
			}

			break

		}

	}
}
