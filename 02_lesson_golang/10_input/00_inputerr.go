// pass_fail сообщает, сдал ли пользователь экзамен.
//Внимание: этот код не будет компилироваться в том виде, в котором он здесь приведен
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input := reader.ReadString('\n')
	fmt.Println(input)
}
