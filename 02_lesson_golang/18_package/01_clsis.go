// програма преобразует температуру в градусах по Фаренгейту в градусы по Цельсию
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}

func main() {
	fmt.Print("Enter a temperature in F:")
	fah, err := getFloat()
	if err != nil {
		log.Fatal(err)
	}
	cel := (fah - 32) * 5 / 9
	fmt.Printf("%0.2f degrees Celsius\n", cel)
}
