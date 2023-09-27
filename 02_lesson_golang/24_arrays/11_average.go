// average вычисляет среднее значение.
package main //Это исполняемая программа, поэтому используем пакет «main».
import "fmt"

func main() {
	numbers := [3]float64{71.8, 56.2, 89.5} //Литерал массива используется для создания массива с тремя числами, для которых вычисляется среднее.
	var sum float64 = 0                     //Объявляем переменную float64 для хранения суммы трех чисел.
	for _, number := range numbers {        //Перебор всех чисел в массиве.
		sum += number //Текущее число бавляется к сумме
	}
	fmt.Println(sum)
}
