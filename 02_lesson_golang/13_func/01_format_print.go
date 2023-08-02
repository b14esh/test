package main

import "fmt"

func main() {
	fmt.Println("About one-third:", 1.0/3.0)
	fmt.Printf("About one-third: %0.2f\n", 1.0/3.0)
        resultString := fmt.Sprintf("About one-third: %0.2f\n", 1.0/3.0)
        fmt.Printf(resultString)
}
