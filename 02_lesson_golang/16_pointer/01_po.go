package main

import "fmt"

func main() {
	myInt := 4 //Значение 
	fmt.Println(myInt) //Вывести значение
	myIntPointer := &myInt //Создаем указатель
	*myIntPointer = 8 //Новое значение присваивается переменной, на которую ссылается указатель (myInt).
	fmt.Println(*myIntPointer) //Выводится значение переменной, на которую ссылается указатель.
	fmt.Println(myInt) //Значение переменной выводится напрямую.
}

/*
4 Исходное значение myInt.
8 Результат обновле-ния *myIntPointer.
8 Обновление значения myInt (то же, что *myIntPointer).
*/
