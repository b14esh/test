package main

import "fmt"

// Директивы package и imports, определения типов пропущены

type part struct {
	description string
	count       int
}

func minimumOrder(description string) part { //Объявляется одно возвращаемое значение типа «part».
	var p part //Создание нового значения «part».
	p.description = description
	p.count = 100
	return p //Функция возвращает тип «part».
}

func main() {
	p := minimumOrder("Hex bolts") //Вызывает minimumOrder. Для сохранения возвращаемого значения «part» использу-ется короткое определение переменной.
	fmt.Println(p.description, p.count)
}
