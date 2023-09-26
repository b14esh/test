package main

import "fmt"

func main() {
	var notes [7]string = [7]string{"do", "re", "mi", "fa", "so", "la", "ti"} //Присваивание значений в формелитерала массива
	fmt.Println(notes[3], notes[6], notes[0])

	var primes [5]int = [5]int{2, 3, 5, 7, 11} // Присваивание значений в форме литерала массива.
	fmt.Println(primes[0], primes[2], primes[4])

}
