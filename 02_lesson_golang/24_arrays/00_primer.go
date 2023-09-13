package main

import (
	"fmt"
	"time"
)

func main() {

	//Пример массив нот:
	var notes [7]string
	notes[0] = "do"
	notes[1] = "re"
	notes[2] = "mi"
	fmt.Println(notes[0])
	fmt.Println(notes[1])

	//Массив целых чисел:
	var primes [5]int
	primes[0] = 2
	primes[1] = 3
	fmt.Println(primes[0])

	//Массив значений time.Time:
	var dates [3]time.Time
	dates[0] = time.Unix(1257894000, 0)
	dates[1] = time.Unix(1447920000, 0)
	dates[2] = time.Unix(1508632200, 0)
	fmt.Println(dates[1])

}
