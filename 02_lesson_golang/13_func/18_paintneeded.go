package main

import "fmt"

func paintNeeded(width float64, height float64) float64 { //объявляет, что paintNeeded возвращает число с плавающей точкой
	area := width * height
	return area / 10.0 //Функция возвращает расход краски вместо того, чтобы выводить его.
}
func main() {
	var amount, total float64                   //Объявляем переменные для хранения расхода краски для текущей сметы, а также для общего расхода по всем сменам.
	amount = paintNeeded(4.2, 3.0)              //Вызываем paintNeeded и сохраняем возвращаемое значение
	fmt.Printf("%0.2f liters needed\n", amount) //Выводим расход для первой сметы
	total += amount                             //Прибавляем расход для текущий смены к total
	amount = paintNeeded(5.2, 3.5)
	fmt.Printf("%0.2f liters\n", total)
	total += amount
	fmt.Printf("Total: %0.2f litters\n", total) //Выводим общий расход по всем сменам
}
