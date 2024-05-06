package main

import "fmt"

func main() {

	// Срез как встроенный тип
	//type slice struct {
	//	Lenght int
	//	Capacity int
	//	ZeroElement *byte
	//}

	//Длина и емкасть slice
	worldSlice := []string{"one", "two", "three"}
	fmt.Println("slice:", worldSlice, "Lenght", len(worldSlice), "Capacity:", cap(worldSlice))
	// Cap - показывает сколько элементов можно запихнуть без выдиления памяти
	worldSlice = append(worldSlice, "four")
	fmt.Println("slice:", worldSlice, "Lenght", len(worldSlice), "Capacity:", cap(worldSlice))
	// Capasity (cap) = или емкасть слайса - это значение показывающие СКОЛЬКО ЭЛЕМЕНТОВ В ПРИНЦИПЕ можно добавить в срез БЕЗ ВЫДИЛЕНИЯ ПАМЯТИ ПОД МАССИВ

	// Допустим у нас есть срез на 3 элемента (иницилизировали без явного указания массива)
	// Компилятор при создания среза слздаст массив ровно на три элемента, после этого вернул адресс где находится этот массив
	// Срез запомнил этот аддресс и теперь ссылается на него
	// Потом
	// Начинаем деформировать слайс(увеличили длину / увеличили кол-во элементов)
	// Проблема - в ниже стоящем массиве (на который ссылается слайс) всего 3 ячейки. Что делать?
	// Компелятор ищет в памяти место для массива размер 3*2 (в общем случае n*2, где n - первоначальный размер)
	// старые 3 элемента на свои позиции. На 4-ую позицию мы добавили элемент
	// После этого компилятор возвращает нашему слайсу новый адрес в памяти, где находится массив 6 элементов

	// Емкасть всегда будет изменятся как n*2

	numerics := []int{1, 2}

	for i := 0; i < 200; i++ {
		if i%5 == 0 { //Если i кратно 5 то выводим следующие
			fmt.Println("Current len:", len(numerics), "Current cap:", cap(numerics))
		}
		numerics = append(numerics, i)
	}
	//Важно: после выделения под новый массив, ссылки со старым будут перенесены в новый
}