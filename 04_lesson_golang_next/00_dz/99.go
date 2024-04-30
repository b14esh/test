package main

import (
	"fmt"
	"strings"
)

func menu() {
	fmt.Println("Add Дата Событие - добавление события")
	fmt.Println("Del Дата Событие - удаление события")
	fmt.Println("Del Дата - удаление всех событий за конкретную дату")
	fmt.Println("Find Дата - поиск событий за конкретную дату")
	fmt.Println("Print - печать всех событий за все даты")
	fmt.Println("StartApp - команда, символизирующая начало работы с трекером")
	fmt.Println("Quit или q - команда завершения работы трекера.")
}

func main() {

	var incom string

outline:
	for {
		fmt.Scanln(&incom)
		switch {
		case strings.Contains(incom, "Quit") || strings.Contains(incom, "q"):
			break outline
		default:
			menu()
		}

	}

}
