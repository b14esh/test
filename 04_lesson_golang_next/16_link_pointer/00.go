package main

import "fmt"

//1. Указатели - перемення, хранящая в качесте хначения - адрес в памяти другой переменной

func main() {

	//2. Определение указателя на что-то
	var variable int = 30
	var pointer *int = &variable //&... - операция взятия адреса в памяти
	//выше у нас содан указатель на переменную variable
	//в pointer лежит 0xc00000a0a8 - это место в памяти, где хранится int  значение 30

	fmt.Printf("Type of pointer %T\n", pointer)
	fmt.Printf("Type of pointer %v\n", pointer)

	//3. А какое нулевое значение для указателя
	var zerPointer *int // zeroValue имеет значение nil (это указатель в никуда)
	fmt.Printf("Type %T and %v\n", zerPointer, zerPointer)
	if zerPointer == nil {
		zerPointer = &variable
		fmt.Printf("After initilization type %T and balue %v\n", zerPointer, zerPointer)
	}

	// Указания поинтера на поинтер // ТАК ЛУЧШЕ НЕ ДЕЛАТЬ
	//zerPointerPointer := zerPointer // ТАК ЛУЧШЕ НЕ ДЕЛАТЬ
	//fmt.Println(zerPointerPointer)  // ТАК ЛУЧШЕ НЕ ДЕЛАТЬ

	//4. Разименования указателя (получения значения): *pointer возращает значение хранимое по адрессу

	var numericValue int = 32
	var ponterToNumeric *int = &numericValue
	//ponterToNumeric := &numericValue
	fmt.Printf("value in numericValue is %v\nAddress is %v\n", *ponterToNumeric, ponterToNumeric)

	//5. Создание указателей на нуливые значения типов
	//var zeroVar int
	//var zeroPoint *int = &zeroVar
	zeroPoint := new(int) // Создает под капотом zeroValue для int, и возращает адрес, где этот 0 хранится
	fmt.Printf("Value in *zeroPoint: %v Address is: %v\n", *zeroPoint, zeroPoint)

	//6. Изменение значения хранимомо по адресу через указатель
	zeroPointerToInt := new(int)
	fmt.Printf("Value in %v and value in zeroPointerToInt is:%v\n", zeroPointerToInt, *zeroPointerToInt)
	*zeroPointerToInt += 40
	fmt.Printf("Value in %v and new value in zeroPointerToInt is:%v\n", zeroPointerToInt, *zeroPointerToInt)

	b := 345
	a := &b
	*a++
	c := &b
	*c += 100
	fmt.Println(b)

	//7. Указательная арифметика // ОТСУСТВУЕТ ПОЛНОСТЬЮ
	//cpp_like := "cpp"
	//cpp_pointer := &cpp_like
	//cpp_pointer++
	//\00.go:59:2: invalid operation: cpp_pointer++ (non-numeric type *string)
	//В GO указательной арифметики НЕТ! и это хорошо

}
