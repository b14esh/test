package main

import (
	"fmt"
	"os"
)

var (
	data                         = []byte{}
	file_name             string = ""
	write_world                  = []byte("Hello World")
	write_world_file_name string = "hello.bin"
)

// func write files
func write_file(date []byte, file_name string) {
	file, err := os.Create(file_name)
	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	file.Write(data)

	fmt.Println("Done.")
}

func main() {
	write_file(write_world, write_world_file_name)
}
