package main

import "fmt"


type subscriber struct {
	name        string
	rate        float64
	active      bool
	homeaddress address
}

type employee struct {
	same        string
	salary      float64
	homeaddress address
}

type address struct {
	street     string
	city       string
	state      string
	postalcode string
}

func main() {
	subscriber := subscriber{}
	fmt.Printf("%#v\n", subscriber.homeaddress)

}

