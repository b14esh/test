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

import "fmt"

func main() {
	var str string

	//	var lenstring int

	//outline:
	//	for {
	//		fmt.Scan(&str)
	//		lenstring = utf8.RuneCountInString(str)
	//		//fmt.Println(lenstring)
	//		if lenstring >= 0 {
	//			fmt.Printf("%s\n", str)
	//		} else {
	//			break outline
	//		}
	//
	//	}

	//	for {
	//		fmt.Scan(&str)
	//		if strings.Contains(str, "") {
	//			return
	//		}
	//		fmt.Printf("%s\n", str)
	//	}

	//LOL
	//command-line-arguments: open C:\Users\b14esh\AppData\Local\Temp\go-build4090313753\b001\exe\18.exe: Operation did not complete successfully because the file contains a virus or potentially unwanted software.
	//	for {
	//		if len(str) == 0 {
	//			return
	//		} else {
	//			fmt.Printf("%s\n", str)
	//		}

	//	}

	for {
		if len(str) > 0 {
			fmt.Printf("%s\n", str)

		} else {
			return
		}

	}

}
