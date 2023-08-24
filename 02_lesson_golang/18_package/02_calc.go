package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getPislo() (int, error){
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	xxx, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}
	return xxx, err
}

func main() {
	fmt.Print("Inter pislo: ")
	calc, err := getPislo()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(calc)
}
