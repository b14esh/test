package main

import (
	"fmt"
	"os"
)

var (
	text                         = ""
	file_name             string = ""
	write_world                  = "Hello World\n ccxxxc \n adaqwerrt \n aaaa "
	write_world_file_name string = "hello.bin"
)

// func write files
func write_file(text string, file_name string) {
	file, err := os.Create(file_name)
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	file.WriteString(text)

	fmt.Println("Done.")
}

func main() {
	write_file(write_world, write_world_file_name)
}
