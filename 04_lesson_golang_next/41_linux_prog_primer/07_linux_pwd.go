// Вывод пути к текущей рабочей директории
// linux
// pwd
package main

import (
	"fmt"
	"os"
)

func pwd() {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	fmt.Println(path)
}

func main() {
	pwd()
}
