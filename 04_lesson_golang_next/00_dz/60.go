package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Читаем файл построчно, затем собираем в массив строк и снова записываем обратно в файл
// Концом строки считаем \n
// Нормально работать будет только в linux
// В windows конец строки \r\n
func edit_string_for_file(stringNeeded string, stringToReplace string, filePath string) {
	text := ""
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		text = scanner.Text()
	}
	if scanner.Text() == stringNeeded {
		text = stringToReplace
	}

	lines = append(lines, text)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(filePath, []byte(strings.Join(lines, "\n")), 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	x := "Строка которую заменить"
	y := "На что заменить"
	//filePath = "/tmp/test.txt"
	i := "test.txt"
	edit_string_for_file(x, y, i)
}
