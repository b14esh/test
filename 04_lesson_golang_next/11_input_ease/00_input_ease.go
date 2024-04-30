package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var name string
	//fmt.Scan(&name) // работает плохо
	// вместо него есть так называемые буферезированые интерфейсы ввода и вывода //
	// Например bufio

	input := bufio.NewScanner(os.Stdin)
	if input.Scan() { //Команда для захвата ввода и сохранения в буфер (захват идет до оканчания сивола конца строки \n)
		name = input.Text() //Команда возращения элементов, помещенных в буфер (на самом деле отдаст строку string)
	}
	fmt.Println(name)

	fmt.Println("For loop started:")
	for {
		if input.Scan() {
			result := input.Text()
			if result == "" {
				break
			}
			fmt.Println("Input string is:", result)
		}
	}
}
