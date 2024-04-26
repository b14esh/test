//Вакуум

//Пользователь вводит строки одну за другой (каждая с новой строки) до тех пор, пока не будет введена пустая строка.
//Программа должна выводить введенные строки, пока не встретилась пустая.

//Ввод	Вывод
//Май     Май
//лайф    лайф
//май     май

//рулез

//Горбатого    Горбатого
//могила       могила

//исправит

package main

import (
	"fmt"
	"strings"
)

func main() {

	var str string
	var snull string = ""
out:
	for {

		fmt.Scan(&str)
		if strings.Contains(str, "1234") {
			break out
		}
		if strings.Contains(str, "\n") {
			break out

			//if strings.HasPrefix(str, thirdPrefix) {
			//	break out
			//}

			//if strings.CutPrefix(str, thirdPrefix) {
			//	break out
			//}
		}
		if strings.Count(str, snull) == -1 {
			break out
		} else {
			fmt.Println(str)
		}

	}

}
