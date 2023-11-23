package main

import "fmt"

type Number int

func (n Number) Add(otherNumber int) {
	fmt.Println(n, "plus", otherNumber, "is", int(n)+otherNumber)
}

func (n Number) Subtract(otherNumber int) {
	fmt.Println(n, "minus", otherNumber, "is", int(n)-otherNumber)
}

func main() {
	ten := Number(10) //Целое число преобразуется в Number.
	ten.Add(4)
	ten.Subtract(5)
	four := Number(4) //Целое число преобразуется в Number.
	four.Add(3)
	four.Subtract(2)
}
