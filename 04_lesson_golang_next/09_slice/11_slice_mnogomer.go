package main

import "fmt"

func main() {
	//Многомерные срезы
	mSlice := [][]int{
		{1, 2},
		{1, 2, 3, 4},
		{10, 20, 30},
		{},
	}

	fmt.Println("mSlice", mSlice)
}
