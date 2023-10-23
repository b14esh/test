package main

import "fmt"

/*
import (
"fmt"
"geo"
)


package geo
type Coordinates struct {
float64
float64
}
*/

type coordinates struct {
	latitude  float64
	longitude float64
}

func main() {
	location := coordinates{latitude: 37.42, longitude: -122.08}
	fmt.Println("Latitude:", location.latitude)
	fmt.Println("Longitude:", location.longitude)
	//fmt.Println("Latitude:", location.Latitude)
	//fmt.Println("Longitude:", location.Longitude)
}
