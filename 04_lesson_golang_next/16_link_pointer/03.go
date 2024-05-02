package main

import "fmt"

//1.Указатели на массивы и почему так делать не нада
//НО ТАК ЛУЧШЕ НЕ ДЕЛАТЬ, НУЖНО ВМЕСТО МАССИВОВ ИСПОЛЬЗОВАТЬ СЛАЙС
func mutation(arr *[3]int) {
	//(*arr)[1] = 909
	//(*arr)[2] = 100000
	//сделаем тоже самое  покороче // Go сам разоменует указатель на массив
	arr[1] = 909
	arr[2] = 10000

}

func main() {
	arr := [3]int{1, 2, 3}
	fmt.Println("Arr befor mutation:", arr)
	mutation(&arr)
	fmt.Println("Arr after mutation", arr)
}
