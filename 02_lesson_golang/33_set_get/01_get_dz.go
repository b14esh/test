package main

import (
	"errors"
	"fmt"
	"log"
)

//coordinates
type Coordinates struct {
        latitude  float64 //Поля должны быть не экспортируемыми
        longitude float64 //Поля должны быть не экспортируемыми
}

func (c *Coordinates) Latitude() float64 {
        return c.latitude
}

func (c *Coordinates) Longitude() float64 {
        return c.longitude
}

func (c *Coordinates) SetLatitude(latitude float64) error {
        if latitude < -90 || latitude > 90 {
                return errors.New("invalid latitude")
        }
        c.latitude = latitude
        return nil
}

func (c *Coordinates) SetLongitude(longitude float64) error {
        if longitude < -180 || longitude > 180 {
                return errors.New("invalid longitude")
        }
        c.longitude = longitude
        return nil
}



//landmark
type Landmark struct {
	name string
	Coordinates
}

func (l *Landmark) Name() string {
	return l.name
}

func (l *Landmark) SetName(name string) error{
	if name == "" {
		return errors.New("invalid name")
	}
	l.name = name
	return nil
}

func main() {
	location := Landmark{}
	err := location.SetName("The Googleplex")
	if err != nil {
		log.Fatal(err)
	}
	err = location.SetLatitude(37.42)
	if err != nil {
		log.Fatal(err)
	}
	err = location.SetLongitude(-122.08)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(location.Name())
	fmt.Println(location.Latitude())
	fmt.Println(location.Longitude())
}
