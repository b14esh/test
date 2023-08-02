//guess - игра, в которой игрок должен отгадать случайное число.
package main

import (
	"bufio"     // пакет отвечает за input keyboard
	"fmt"       // вывод в консоль
	"log"       // работает с ошибками err
	"math/rand" // отвечает за генерацию чисел
	"os"        // использование возможностей ОС
	"strconv"   // работа со строками, преобразование
	"strings"   // работа со строками
	"time"      // работает со временем, используется
)

func main() {
	seconds := time.Now().Unix() //Получает текущую дату и время в виде целого числа.
	rand.Seed(seconds)           // Инициализируем генератор случайных чисел.
	target := rand.Intn(100) + 1 // Генерируем целое число 1 от 100
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")

	reader := bufio.NewReader(os.Stdin) // Создаем bufio.Reader для чтения ввода с клавиатуры.
	success := false                    // Настроить программу, чтобы по умолчанию выводилось сообщение о проигрыше.

	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You have", 10-guesses, "guesses left.")
		fmt.Print("Make a guess: ")           // Запрашиваем число.
		input, err := reader.ReadString('\n') // Прочитать данные, введенные пользователем до нажатия Enter.
		if err != nil {
			// Если произошла ошибка программа выводит сообщение и завершается
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)  // Удаляем символ новой строки.
		guess, err := strconv.Atoi(input) // Введенная строка преобразуется в целое число
		if err != nil {
			// Если произошла ошибка программа выводит сообщение и завершается
			log.Fatal(err)
		}
		if guess < target {
			// если введенное значение меньше загаданного, сообщить об этом.
			fmt.Println("Oops. Your guess was LOW")
		} else if guess > target {
			// Если введено сообщение больше загаданного, сообщить об этом
			fmt.Println("Oops.Your guss was HIGH")
		} else {
			// В противном случае введенное значение должно быть правильным...
			success = true //Предотвращает вывод сообщения о проигрыше.
			fmt.Println("Good job! You guessed it")
			break //Выход из цикла
		}
	}
	if !success {
		// Если переменная «success» равна false, сообщить игроку загаданное число.
		fmt.Println("Sorry, you diddn't guess my number. It was:", target)
	}
}
