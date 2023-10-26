//Если вы хотите сложить значение типа Liters со значением типа Gallons,
//сначала необходимо преобразовать один тип и привести его в соответствие с другим.
package main

import "fmt"

type Liters float64
type Gallons float64

func main() {

	fmt.Println(Liters(1.2) + Gallons(3.4))
	fmt.Println(Gallons(1.2) == Liters(1.2))

}
