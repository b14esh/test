package main

import (
	"fmt"
	"unicode/utf8"
)

// 1. Описание того что должен уметь претиндент
type BigWorld interface {
	IsBig() bool
}

// 2. Наш претиндент которого нужно научить делать IsBig() bool
type MySuperString string

// 3. Реализация IsBig
func (mss MySuperString) IsBig() bool {
	if utf8.RuneCountInString(string(mss)) > 10 {
		return true
	}
	return false
}

func main() {

	sample := MySuperString("saffjndsijnfinfinwdnfw")
	var interfaceSample BigWorld //Обьявление переменной, типа интерфеса BigWorld
	interfaceSample = sample     // присваивание значения для переменной тип BigWorld возможно,
	//потому что sample (типа MySuperString) удовлетворяет интерфейсу BigWord
	fmt.Println("isBig?", interfaceSample.IsBig())

	//4. Попытка присвоить значения типа неудовлетворяещему интерфейсу
	//	var interfaceBadSample BigWorld
	//	interfaceBadSample = "asdaf" //тип стринг не имеет реализации метода isBig, поэтому не удовлетворяет интерфейсу

}
