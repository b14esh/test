package main

import (
	"bufio"
	"os"
)

// Вывод строки текста
func echo(text ...string) {

	/*Тут мы сначала, для подготовки выходного значения, создаём байтовый срез.
	Потом перебираем входные строки, присоединяя их байты к срезу.
	А перед обработкой строк (за исключением первой) мы присоединяем к срезу один пробельный символ.
	В итоге мы добавляем к срезу символ новой строки и выводим то, что получилось, в stdout.
	*/
	var textBytes []byte

	for i, t := range text {
		if i > 0 {
			textBytes = append(textBytes, ' ')
		}

		textBytes = append(textBytes, t...)
	}

	textBytes = append(textBytes, '\n')

	if _, err := os.Stdout.Write(textBytes); err != nil {
		panic(err)
	}
}

func main() {

	input := bufio.NewScanner(os.Stdin)

	input.Scan()
	result := input.Text()
	if result == "" {
		return
	}

	echo(result)
}
