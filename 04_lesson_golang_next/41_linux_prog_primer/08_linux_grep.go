//Вывод строк файла, соответствующих заданному шаблону
//grep local /etc/hosts

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func grep(pattern string, fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()

		if strings.Contains(text, pattern) {
			fmt.Println(text)
		} else {
			fmt.Println("NO")

		}
	}
}

func main() {
	grep("local", "hosts")
}
