package main

import "fmt"

func main() {
	fmt.Print("Enter a grade:")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	if input >= 60 {
		status := "passing"
	} else {
		status := "failing"
	}

}
