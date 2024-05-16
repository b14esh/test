package main

import "fmt"

func main() {

	fmt.Println("Numerecs.float")
	//Numerics. Float
	//float32, float64

	floatfirst, floatsecond := 5.67, 12.64

	fmt.Printf("type of a %T, type of b %T \n", floatfirst, floatsecond)

	// - + * / (не льзя взять остаток от деления %)

	sum := floatfirst + floatsecond
	sub := floatfirst - floatsecond
	suc := floatfirst * floatsecond
	sud := floatfirst / floatsecond

	fmt.Println("sum =", sum, "sub =", sub, "suc", suc, "sud", sud)
	//Для float чисел не рекомендуется использовать функцию Println !!!
	//Для float чисел не рекомендуется использовать функцию Println !!!

	// Используйте Printf
	fmt.Printf("sum = %f, sub = %f \n", sum, sub)

	// Используйте Printf и указывайте кол-во знаков после запятой %.4f, %.3f
	fmt.Printf("suc = %.3f, sud = %.3f \n", suc, sud)

}
