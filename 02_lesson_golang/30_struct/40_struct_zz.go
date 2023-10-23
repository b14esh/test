package main

import "fmt"

/*

package geo
type Coordinates struct {
Latitude float64
Longitude float64
}
type Landmark struct {
	Name string
	Coordinates
}

package main
import (
"fmt"
"geo"
)
func main() {
location := geo.Landmark{}
location.Name = "The Googleplex"
location.Latitude = 37.42
location.Longitude = -122.08
fmt.Println(location)
}

*/



type Coordinates struct {
	Latitude  float64
	Longitude float64
}
type Landmark struct {
	Name string
	Coordinates
}

func main() {
	location := Landmark{}
	location.Name = "The Googleplex"
	location.Latitude = 37.42
	location.Longitude = -122.08
	fmt.Println(location)
}
