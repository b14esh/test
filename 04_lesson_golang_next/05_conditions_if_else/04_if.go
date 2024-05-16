package main

import (
	"fmt"
)

func main() {

	fmt.Println("Иницилизация в блоке услоного операторая")
	//for ( int i = 0; i < 0; i++){}

	// Хорошо
	// Блок присваивания для if только вот так вот ":="
	//Инициализируемая переменная видна ТОЛЬКО внутри области виидимости условного оператора (в телах if, else if, или else)
	// Но не за его пределами
	if num := 10; num%2 == 0 {
		fmt.Println("EVEN", num)
	} else {
		fmt.Println("ODD", num)
	}

	//Ущербно
	/*
		var age int = 10
		if age > 7 {
			fmt.Println("Go to school")
		} //По факту, сюда подставляется ; компилятором, и дальнейший код уже не имеет связи с предыдущим if
		else {
			fmt.Println("Another case")
		}
	*/

	//НЕ ИДЕОМАТИЧНО
	if width := 100; width > 100 {
		fmt.Println("Width > 100")
	} else {
		fmt.Println("Width <= 100")
	}
	//Странное правило номер 1: в Go стараются избегать блоков ELSE

	//Идеоматичность
	if height := 100; height > 100 {
		fmt.Println("height > 100")
		return //Провоцирут завершение функции main и дальше работа не будет производится
	}
	fmt.Println("Height <= 100")
}
