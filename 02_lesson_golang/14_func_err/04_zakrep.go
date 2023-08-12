package main

import "fmt"

func nl(needstring int) {
	for i := 1; i < needstring; i++ {
		fmt.Print("\n")
	}
}

func nl2() {
	i := 1
	for i <= 3 {
		fmt.Println("i =", i)
		i++
	}
}

func nl3() {
	i := 5
	for i > 0 {
		fmt.Println("i =", i)
		i--
	}
}

func nl4() {
	for i := 10; i >= 7; i-- {
		fmt.Println("i =", i)
	}
}

func summ1(num1 int, num2 int, num3 int) int {
	return num1 + num2 + num3
}

func summ2(num1 int, num2 int) (int, int) {
	var res2 int
	var res1 int
	res1 = num1 + num2
	res2 = num1 * num2
	return res1, res2
}

func main() {
	//comment
	//var age = 13
	//var num = 7.1234
	//var str = "Hello"
	age := 13
	num := 7.123
	str := "hello"
	fmt.Println("\ninteger:", age, "\nfloat64:", num, "\nSting:", str)

	b := 2
	a := b
	res := a + b
	fmt.Println("\nResult a + b  is:", res)

	c := 100
	res1 := c - a*b
	fmt.Println("\nResult c - a*b is:", res1)

	x1 := 10
	x2 := 5
	res2 := x1 % x2
	fmt.Println("\nx1 % x2 \nОстаток при делении:", res2)

	web := "lollolololooooro"
	fmt.Println("\nПосчетаем кол-во символов в строке из переменной web:", len(web))
	fmt.Println("Обьеденить строки:", web+web+" xxxx")

	var num1 float64 = 7.12345678
	fmt.Printf("\nОставляем в дробной части два символа: %.2f ", num1)
	fmt.Printf("\nУзнаем тип переменной num1 : %T \n", num1)

	var isDone bool = true
	fmt.Printf("\nisDone равно:\n\t\t\t %T \n", isDone)

	var isDone1 bool = false
	fmt.Printf("\n%t\n", isDone1)

	nl(2)
	//Условные операторы
	//c2 := 9
	//c2 := 7
	c2 := 200
	if c2 < 5 {
		fmt.Println("c2 > 5 , c2 = ", c2)
	} else if c2 == 5 {
		fmt.Println("c2 == 5 , c2 =", c2)
	} else if (c2 > 5) && (c2 <= 18) {
		fmt.Println(" c2 > 5 и c2 <= 18 , c2 =", c2)
	} else {
		fmt.Println("c2 =", c2)
	}

	nl(11)

	// Switch case
	l1 := 100
	switch l1 {
	case 5:
		fmt.Println("l1 =", l1)
	case 7:
		fmt.Println("l1 =", l1)
	case 10:
		fmt.Println("l1 =", l1)
	case 20:
		fmt.Println("l1 =", l1)
	default:
		fmt.Println("Вы вышли из значений case, l1 = ", l1)
	}

	nl(5)

	nl2()

	nl(5)

	nl3()

	nl(10)

	nl4()

	nl(4)
	// массивы
	var arr [5]int
	arr[0] = 45
	arr[1] = 60
	arr[2] = 100
	arr[3] = 65
	arr[4] = 200
	fmt.Println(arr[2])

	nl(2)
	//Пример перебора массива
	nums2 := [3]float64{4.23, 5.44, 56.12}
	for i, value := range nums2 {
		fmt.Println(value, i)
	}
	nl(2)
	// Карты, maps
	xData := make(map[string]float64)
	xData["xxxx"] = 0.16
	xData["yyyy"] = 0.99
	fmt.Println(xData["xxxx"])

	nl(5)

	fmt.Println(summ1(200, 555, 234))

	fmt.Println(summ2(100, 200))
}
