package main

import "fmt"

func main() {
	// Структуры
	bob := Cats{"BOB", 7, 0.77}
	fmt.Println("Bob age is", bob.age)
	fmt.Println("Bob function", bob.test())

}


type Cats struct {
	name string
	age int 
	happiness float64
}


func (cat *Cats) test() float64 {
	return float64(cat.age) * cat.happiness
}
