package main

import "fmt"

func main() {

	box := map[int]string{1: "a", 2: "sak", 100: "xxx"}

	boxi := map[string]int{"xa": 2, "za": 100, "ddd": 400}

	boxb := map[string]bool{"yes": true, "no": false}

	for k, v := range box {
		fmt.Println(k, v)
	}

	for key, value := range boxi {
		fmt.Println(key, value)
	}

	for x, i := range boxb {
		fmt.Println(x, i)
	}

	mapss := map[int]float64{1: 22.5, 2: 33.1, 4: 41.2}
	key, value := mapss[2]
	fmt.Println(key, value)

	mapsss := map[string]string{"xox": "lox", "box": "xor"}
	k, g := mapsss["box"]
	fmt.Println(k, g)

}
