package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	/*
	   Чтобы присвоить значение переменной, используйте os.Setenv.
	   Для получения значения - os.Getenv.
	   Эта функция вернет пустую строку, если такой переменной нет в окружении.
	*/
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))
	/*
	   Используйте os.Environ, чтобы вывести все существующие переменные
	   и их значения. Эта функция возвращает срез строк в формате КЛЮЧ=значение.
	   Вы можете strings.Split-нуть их, чтобы разделить на ключ и значение.
	   Тут мы выводим все ключи на экран.
	*/
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}
}
