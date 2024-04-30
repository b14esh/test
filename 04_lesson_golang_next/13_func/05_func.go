package main

import (
	"fmt"
)

func main() {
	helloVariadic(10, 20, 30, 40, 50, 60, 70)
	helloVariadic(10, 20)
	helloVariadic()

	someStrings(2, 3, "bob", "alex", "vita")

	sum1 := sumVariadic(10, 20, 30, 34, 44)
	sliceNumber := []int{10, 20, 30, 40}
	sum2 := sumVariadic(sliceNumber...)
	fmt.Println(sum1, sum2)

	//12. Анонимная функция (синтаксис)
	anon := func(a, b int) int {
		return a + b
	}
	fmt.Println(anon(20, 30))

	fmt.Println("BigFunc(10, 20):", bigFunction(10, 20))
}

// 9. Variadic parametrs (Континуальные параметры)
func helloVariadic(a ...int) {
	fmt.Println(a)
	fmt.Printf("value %v and type %T", a, a)
}

// 10. Смешение параметров с variadic
// func someStrings(words ...string, a int, b int) //будет ошибка //  Вариадики нужно добавлять в конец
// 1. Контитульный параметр всегда самый последний (...string )
// 2. Вариадик параметр на всю фунцию ОДИН
func someStrings(a int, b int, words ...string) {
	fmt.Println("Parametr:", a)
	fmt.Println("Parametr:", b)
	var result string
	for _, word := range words {
		result += word
	}
	fmt.Println("result:", result)

}

//11. Передача слайса или использование вариадик параметра

func sumVariadic(nums ...int) int {
	var sum int
	for _, val := range nums {
		sum += val
	}
	return sum
}

// 13. Ананимная функция внутри явной
func bigFunction(aArg, bArg int) int {
	return func(a, b int) int { return a + b + 1 }(aArg, bArg)

}
