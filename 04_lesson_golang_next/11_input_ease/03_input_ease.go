package main

import (
	"fmt"
	"strconv"
)

func main() {

	//Преобразование строкового литерала к чему-нибудь числовому
	numstr := "10"
	numint, _ := strconv.Atoi(numstr) // Atoi - Anything to int (именно int) (Оно там ошибки возращает, но мы пока не усеем с ними работать по этом _)
	fmt.Printf("%d is %T \n", numint, numint)

	numint64, _ := strconv.ParseInt(numstr, 10, 64)
	fmt.Printf("%d is %T\n", numint64, numint64)

	numFloat64, _ := strconv.ParseFloat(numstr, 64) // отдает только 64 но можно без проблем привести к 32
	fmt.Printf("%f is %T\n", numFloat64, numFloat64)
}
