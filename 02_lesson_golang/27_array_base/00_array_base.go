package main

import "fmt"

func main() {

	array1 := [5]string{"a", "b", "c", "d", "e"}
	slice1 := array1[0:3]
	fmt.Println(slice1)

	array2 := [5]string{"f", "g", "h", "i", "j"}
	slice2 := array2[2:5]
	fmt.Println(slice2)

	array3 := [5]string{"a", "b", "c", "d", "e"}
	slice3 := array3[0:3]
	slice4 := array3[2:5]
	fmt.Println(slice3, slice4)

	slice9 := array1[0:3]
	array1[1] = "X"
	fmt.Println(array1)
	fmt.Println(slice9)

	slice10 := array2[2:5]
	slice10[1] = "X"
	fmt.Println(array2)
	fmt.Println(slice10)

	slice11 := array3[0:3]
	slice12 := array3[2:5]
	array3[2] = "X"
	fmt.Println(array3)
	fmt.Println(slice11, slice12)
}
