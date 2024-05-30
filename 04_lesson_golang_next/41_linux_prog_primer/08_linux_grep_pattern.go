// Вывод строк файла, соответствующих заданному шаблону
// grep local /etc/hosts
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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
		matched, _ := regexp.MatchString(pattern, text)

		fmt.Println(matched)
	}
}

func main() {
	grep(`zoc*`, "hosts")
}
