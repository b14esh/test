//Глагол Вывод
//%f     Число с плавающей точкой
//%d     Десятичное целое число
//%s     Строка
//%t     Логическое значение (true или false)
//%v     Произвольное значение (подходящий формат выбирается на основании типа передаваемого значения)
//%#v    Произвольное значение, отформатированное в том виде, в котором оно отображается в коде Go
//%T     Тип переданного значения (int, string и т. п.)
//%%     Знак процента (литерал)

package main

import "fmt"

func main() {

	fmt.Printf("A float: %f\n", 3.1415)
	fmt.Printf("An integer: %d\n", 15)
	fmt.Printf("A string: %s\n", "hello")
	fmt.Printf("A boolean: %t\n", false)
	fmt.Printf("Values: %v %v %v\n", 1.2, "\t", true)
	fmt.Printf("Values: %#v %#v %#v\n", 1.2, "\t", true)
	fmt.Printf("Types: %T %T %T\n", 1.2, "\t", true)
	fmt.Printf("Percent sign: %%\n")
        fmt.Printf("%v %v %v", "", "\t", "\n")
	fmt.Printf("\n")
        fmt.Printf("%#v %#v %#v", "", "\t", "\n")
	fmt.Printf("\n")

}
