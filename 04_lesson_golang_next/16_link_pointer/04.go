package main

import "fmt"

//2. Используйте лучше слайсы вместо ссылок с масивами
// Мапы тоже не имеет смысла добовлять ссылками
// Он также прекрасно как и слайсы работает без ссылок
func mutationSlice(sls []int) {
	sls[1] = 909
	sls[2] = 10000

}

func main() {
	newArr := [3]int{1, 2, 4}
	fmt.Println("Arr befor mutationSlice:", newArr)
	mutationSlice(newArr[:])
	fmt.Println("Arr after mutation", newArr)
}
