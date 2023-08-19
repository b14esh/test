package main

import "fmt"

func az() {
        for x := 0; x < 11; x++ {
                fmt.Println(x)
        }
}
func za() {
        for y := 0; y < 11; y++ {
                fmt.Println(y)
        }
}

func main() {
        az()
        za()
        z := 0
        x := 0
        for i := 0; i< 11; i++ {
                fmt.Println(x, z)
                x++
                z++
        }
}
