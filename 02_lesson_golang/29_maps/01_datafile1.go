// Пакет datafile предназначен для чтения данных из файлов.
package main

import (
	"bufio"
	"fmt"
	"os"
	"log"
)

// GetStrings читает строку из каждой строки данных файла.
func GetStrings(fileName string) ([]string, error) {
	var lines []string
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	err = file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return lines, nil
}

func main() {
	filename := "qotes.txt"
	lines, err := GetStrings(filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(lines)
}
