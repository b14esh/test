package main

import "fmt"

type myStruct struct {
	myField int
}

func main() {
	var value myStruct //Создание значения структуры.
	value.myField = 3
	var pointer *myStruct = &value //Получает указатель на значение структуры.
	fmt.Println(*pointer.myField)  //Пытаемся получить значе-ние структуры, на которое ссылается указатель.
}

//Ошибка
//invalid indirect of
//pointer.myField (type int)
