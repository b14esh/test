package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.Open("data.txt") //Файл данных открыва
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	//Цикл выполняется до того, как будет  достигнут конец  файла, а scanner.Scan вернет false
	for scanner.Scan() { //Читает строку из файла
		fmt.Println(scanner.Text()) //Выводит строку
	}
	// Если при закрытии файла произошла ошибка то сообщить о ней и завершить работу
	err = file.Close() //Закрывает файл для освобождения ресурсов. 
	if err != nil {
		log.Fatal(err)
	}

	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
}
