package main

import "fmt"

func changeParam(val *int) {
	//8.1 Определение функции, принимающией параметр как указатель
	*val += 100
}

func chanheWOPointer(val int) {
	//8.1 функция для примера что будет если не использовать поинт
	// Никак не повлияет
	// Так как все изменения останутся тут
	val += 100
}

func main() {
	//8. Передача указателей в функцию
	sample := 1
	samplePointer := &sample
	fmt.Println("Origin value of sample:", sample)
	changeParam(samplePointer)
	chanheWOPointer(sample) //никак как не повлияет
	fmt.Println("After changing samle is:", sample)

}
