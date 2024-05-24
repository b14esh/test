// Вывод имён файлов, содержащихся в директории, отсортированных в обратном порядке
// ls /bin | sort -r
package main

import (
	"fmt"
	"os"
	"sort"
)

func ls(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	list, err := file.Readdirnames(0)
	if err != nil {
		panic(err)
	}

	sort.Strings(list)

	for i := len(list) - 1; i >= 0; i-- {
		name := list[i]

		fmt.Println(name)
	}
}

func main() {
	ls("/bin")
}
