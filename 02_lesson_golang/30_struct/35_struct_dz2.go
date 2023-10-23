package main

import "fmt"

/*
package magazine
// Код Subscriber
type Address struct {
Street string
City string
State string
PostalCode string
}

package main
import (
"fmt"
"github.com/headfirstgo/magazine"
)
func main() {
var address magazine.Address
address.Street = "123 Oak St"
address.City = "Omaha"
address.State = "NE"
address.PostalCode = "68111"
fmt.Println(address)
}
*/


type Addr struct {
	street     string
	city       string
	state      string
	postalCode string
}

func main() {
	var address Addr
	address.street = "123 Oak St"
	address.city = "Omaha"
	address.state = "NE"
	address.postalCode = "68111"
	fmt.Println(address)
}
