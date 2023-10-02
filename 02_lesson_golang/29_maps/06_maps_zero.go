package main

import "fmt"

func main() {

	numbers := make(map[string]int)
	numbers["I've been assigned"] = 12
	fmt.Printf("%#v\n", numbers["I've been assigned"])
	fmt.Printf("%#v\n", numbers["I haven't been assigned"])

	words := make(map[string]string)
	words["I've been assigned"] = "hi"
	fmt.Printf("%#v\n", words["I've been assigned"])
	fmt.Printf("%#v\n", words["I haven't been assigned"])

	counters := make(map[string]int)
	counters["a"]++
	counters["a"]++
	counters["c"]++
	fmt.Println(counters["a"], counters["b"], counters["c"])
}
