package main

func main() {

	// Массиа и его размер - это две составляющие одного типа (Размер массив - часть типа)
	var aArr [5]int
	var bArr [6]int
	aArr[0] = 100
	bArr = aArr
	// будет ошибка
	// .\07_massiv.go:9:9: cannot use aArr (variable of type [5]int) as [6]int value in assignment

}
