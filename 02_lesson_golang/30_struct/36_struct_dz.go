package main

import "fmt"

/*
type Subscriber struct {
Name string
Rate float64
Active bool
HomeAddress Address
}
type Employee struct {
Name string
Salary float64
HomeAddress Address
}
type Address struct {
// ...
}


main.go
package main
import (
"fmt"
"github.com/headfirstgo/magazine"
)
func main() {
address := magazine.Address{Street: "123 Oak St",
City: "Omaha", State: "NE", PostalCode: "68111"}
subscriber := magazine.Subscriber{Name: "Aman Singh"}
subscriber.HomeAddress = address
fmt.Println(subscriber.HomeAddress)
}


*/


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
	address := address{street: "123 Oak St", city: "Omaha", state: "NE", postalcode: "68111"}
	subscriber := subscriber{name: "Aman Singh"}
	subscriber.homeaddress = address
	fmt.Println(subscriber.homeaddress)
}
