//"Бу'сайз" - дизайнерское агенство, основным полем деятельности которого является разработка логотипов с текстовой начинкой.
//У данного агенства вполне интересный прайс-лист  - за каждый символ (не важно какой) , агенство берет плату в размере 23 коп.

//Вам поручено написать программу, которая выполняет вычислительную операцию подсчета стоимости логотипа в зависимости от количества символов в строке клиента.

//Напишите программу, которая считывает строку и выводит ее предполагаемую цену в следующем формате:

//Y коп. - если цена не дотягивает до рубля.
//X р. Y коп. - если цена превышает 1 рубль.

//Привет?_Вацап?_Кек_:)     4 р. 83 коп.

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	var (
		cost     float64 = 23
		worlds   string
		lenworld int
		fullcost float64
		obrezcop float64
	)

	//worlds = "Привет?_Вацап?_Кек_:)"
	fmt.Scan(&worlds)
	lenworld = utf8.RuneCountInString(worlds)
	fullcost = cost * float64(lenworld)

	if fullcost < 100 {
		fmt.Printf("%0.0f коп.", fullcost)
		return
	}

	i := fullcost / 100
	obrezcop = (i - float64(int(i))) * 100

	//fmt.Println(i)
	//fmt.Println(obrezcop)

	fmt.Printf("%d р. %0.0f коп.", int(i), obrezcop)

}
