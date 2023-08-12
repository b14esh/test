package main

import "fmt"

//структура
type Cars struct {
	model  string
	age    int
	callse float64
}

func main() {
	mersc := Cars{"beh", 20, 0.13}
	fmt.Println("Info for mersc:", mersc.model)
	fmt.Println("mersc is xren:", mersc.hlop())

}


func (gogs *Cars) hlop() float64 {
	return float64(gogs.age) * gogs.callse
}
