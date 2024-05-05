package main

import "fmt"

type Describer interface {
	Describer()
}

type Student struct {
	name string
	age  int
}

func (std Student) Describer() {
	fmt.Printf("%s and %d  y.o\n", std.name, std.age)
}

func TypeFinder(i interface{}) {
	switch v := i.(type) { //Присвоем значение "v" лежащие под предпологаемым интерфейсом
	case string:
		fmt.Println("This is string")
	case int:
		fmt.Println("this is int")
	case Describer: // Вывод с типом интерфейса можно сравиватся как с типажем, как и с другим типом языка
		// это как раз и говорит о том, что интерфейсы полу обстрактные типы.
		fmt.Println("I'm Describer")
		v.Describer()
	default:
		fmt.Println("Unknow type")
	}
}

func buba(a int, b int, x int) int {
	c := a + b + x
	return c
}

func changeParam(val *int) {
	//8.1 Определение функции, принимающией параметр как указатель
	*val += 100
}


func main() {
	a := 10
	b := 10
	x := 10
	fmt.Println(buba(a, b, x))

	sample := 1
	//samplePointer := &sample
	fmt.Println("Origin value of sample:", sample)
	//changeParam(samplePointer)
	changeParam(&sample)
	fmt.Println("After changing samle is:", sample)

	
	zv :=  


	std := Student{"Alex", 23}
	TypeFinder(10)
	TypeFinder("Hello")
	TypeFinder(nil)
	TypeFinder(std)

}
