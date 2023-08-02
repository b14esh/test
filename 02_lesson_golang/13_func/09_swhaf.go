/*Допустим, вы хотите вычислить, сколько
краски потребуется для покрытия нескольких стен.
Производитель указывает, что одного литра краски хватает на 10 квадратных метров.
Следовательно, для определения количества литров необходимо вычислить площадь каждой стены,
умножив ее ширину (в метрах) на высоту, а затем разделить результат на 10.*/

package main

import "fmt"

/*func main() {
	var width, height, area float64
	//Вычислить расход краски для первой стены.
	width = 4.2
	height = 3.0
	area = width * height //Вычисляем площадь стены.
	//fmt.Println(area/10.0, "liters needed") //Вычисляем, сколько краски понадобится для этой площади.
	fmt.Printf("%.2f liters needed\n", area/10.0) //Значение форматируется и вставляется в строку.
	//Сделать то же для этой площади. самое для второй стены.
	width = 5.2
	height = 3.5
	area = width * height
	//fmt.Println(area/10.0, "liters needed")     //Вычисляем, сколько краски понадобится для этой площади.
	fmt.Printf("%.2f liters needed\n", area/10.0) //Значение форматируется и вставляется в строку.

}
*/

func paintNeeded(width float64, height float64) {
	area := width * height
	fmt.Printf("%.2f liters needed\n", area/10.0)
}

func main() {
	paintNeeded(4.2, 3.0)
	paintNeeded(5.2, 3.5)
	paintNeeded(5.0, 3.3)
}
