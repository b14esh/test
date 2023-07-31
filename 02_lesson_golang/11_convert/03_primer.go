package main
import (
        "bufio"
        "log"
        "os"
        "strconv"
        "strings"
        "fmt"
)
func main() {
        fmt.Print("Enter a grage: ")
        reader := bufio.NewReader(os.Stdin)
        input, err := reader.ReadString('\n')
        if err != nil {
                log.Fatal(err)
        }
        input = strings.TrimSpace(input)
        grade, err := strconv.ParseFloat(input, 64)
        if err != nil {
                log.Fatal(err)
        }
        var status string  // вот решение
        if grade >= 60 {
                status = "passing" // так как перемену мы задали выше тут мы уже используем присвоение
        } else {
                status = "failing" // так как перемену мы задали выше тут мы уже используем присвоение
        }
        fmt.Println("A grade of", grade, "is", status)
}
