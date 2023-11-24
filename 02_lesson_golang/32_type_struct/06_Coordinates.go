/*
package geo

type Coordinates struct {
	Latitude float64
	Longtitude float64
}

func (c ) SetLatitude( float64) {
	= latitude
}

func (c )  SetLongitude( float64) {
	= longitude
}



package main

import (
	"fmt"
	"geo"
)

func main() {
	coordinates := geo.Coordinates{}
	coordinates.SetLatitude(37.42)
	coordinates.SetLongitude(-122.08)
	fmt.Println(coordinates)
}
*/

package main

import "fmt"

type Coordinates struct {
	Latitude   float64
	Longtitude float64
}

func (c *Coordinates) SetLatitude(latitude float64) {
	c.Latitude = latitude
}

func (c *Coordinates) SetLongitude(longtitude float64) {
	c.Longtitude = longtitude
}

func main() {
	coordinates := Coordinates{}
	coordinates.SetLatitude(37.42)
	coordinates.SetLongitude(-122.08)
	fmt.Println(coordinates)
}
