package main

import (
	"fmt"
	"math"
)

var PORT = "8000"

//1. Константы это не изменяемые переменные
//1) которые служат для более строго понимания кода
//2) Для того что бы случайно не поменять значения (Предпологается что згачения константы не изменно)
//3) Для удобных преобразований

const (
	//Правила именования константы, стараются писать капсом, а так в GO нет общего правила именования констант
	MAIN_PORT = "8001"
)

func main() {
	//2. Обьявление одной константы
	const a = 10
	fmt.Println(a)

	//3. Обьявление блока констант
	const (
		ipAdress = "127.127.0.3"
		port     = "8000"
		dbname   = "postgress"
	)

	fmt.Println("ipadress value", ipAdress)
	fmt.Println("Cheack port is valid", checkPortisValid())

	//4. Константу никак нельзя поменять
	//cons b = 200
	//b = 30

	//5. Значение констант должны быть известны  на момент компиляции
	var sqrt = math.Sqrt(25) //Квадратный корень из 25 // Нет ошибки
	//const sqrt = math.Sqrt(25) // Ошибка
	//  math.Sqrt(25) (value of type float64) is not constant
	// Нельзя присвоить в констату что ли-бо является результатом вызова функции, метод и т.д
	fmt.Println("var sqrt:", sqrt)

	//6. типезированые и нетипезированые константы
	const ADMIN_EMAIL string = "admin@admin.com" // Указание типа - это повышение читабельности кода
	fmt.Println(ADMIN_EMAIL)

	//7. Нетипизированые константы и их профит
	// При использовании не типизированных констант мы разрешаем компелятору использовать не явное предстовления типов в переменные
	const NUMERIC = 10
	var numeric = NUMERIC
	var numint8 int8 = NUMERIC
	var numint16 int16 = NUMERIC
	var numint32 int32 = NUMERIC
	var numint64 int64 = NUMERIC
	var numfloat32 float32 = NUMERIC
	var numfloat64 float64 = NUMERIC
	var numcomplex64 complex64 = NUMERIC

	fmt.Printf("Value %v and Type %T\n", numeric, numeric)
	fmt.Printf("Value %v and Type %T\n", numint8, numint8)
	fmt.Printf("Value %v and Type %T\n", numint16, numint16)
	fmt.Printf("Value %v and Type %T\n", numint32, numint32)
	fmt.Printf("Value %v and Type %T\n", numint64, numint64)
	fmt.Printf("Value %v and Type %T\n", numfloat32, numfloat32)
	fmt.Printf("Value %v and Type %T\n", numfloat64, numfloat64)
	fmt.Printf("Value %v and Type %T\n", numcomplex64, numcomplex64)

	//8. Константы зашиваются в RUNTIME момент копиляции и не выбрасываются до ее окончании

}

func checkPortisValid() bool {

	if MAIN_PORT == "8001" {
		return true
	}
	return false
}
