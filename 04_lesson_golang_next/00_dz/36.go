//На вход программе подается строка.
//Если из первого и последнего символов этой строки можно собрать слово "да" ("Да", "дА", "ДА", "да") - то программа выводит "СОГЛАСЕН"
//и "НЕ СОГЛАСЕН" в противном случае.

// Ввод
// дома
// Вывод
// СОГЛАСЕН
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	input := bufio.NewScanner(os.Stdin)

	for {
		if input.Scan() {
			result := input.Text()
			if result == "" {
				break
			}

			fullruneSlice := []rune(result) // преобразовываем  result в руны
			//fmt.Printf("%c \n", fullruneSlice[0])
			//fmt.Printf("%c \n ", fullruneSlice[len(fullruneSlice)-1])

			// немного упращаем дальнейший перебор в if
			a := fullruneSlice[0]                    // первая буква
			b := fullruneSlice[len(fullruneSlice)-1] // последняя буква

			if a == 'д' && b == 'а' || a == 'Д' && b == 'А' || a == 'д' && b == 'А' || a == 'Д' && b == 'а' {
				fmt.Println("СОГЛАСЕН")
				break
			} else {
				fmt.Println("НЕ СОГЛАСЕН")
				break
			}

		}
	}

}
