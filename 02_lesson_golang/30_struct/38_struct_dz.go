package main

import "fmt"

type Subscriber struct {
	Name        string
	Rate        float64
	Active      bool
	HomeAddress Address
}

type Employee struct {
	Name        string
	Salary      float64
	HomeAddress Address
}

type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
}

func main() {
	subscriber := Subscriber{Name: "Aman Singh"}
	subscriber.HomeAddress.Street = "123 Oak St"
	subscriber.HomeAddress.City = "Omaha"
	subscriber.HomeAddress.State = "NE"
	subscriber.HomeAddress.PostalCode = "68111"
	fmt.Println("Subscriber Name:", subscriber.Name)
	fmt.Println("Street:", subscriber.HomeAddress.Street)
	fmt.Println("City:", subscriber.HomeAddress.City)
	fmt.Println("State:", subscriber.HomeAddress.State)
	fmt.Println("Postal Code:", subscriber.HomeAddress.PostalCode)
	employee := Employee{Name: "Joy Carr"}
	employee.HomeAddress.Street = "456 Elm St"
	employee.HomeAddress.City = "Portland"
	employee.HomeAddress.State = "OR"
	employee.HomeAddress.PostalCode = "97222"
	fmt.Println("Employee Name:", employee.Name)
	fmt.Println("Street:", employee.HomeAddress.Street)
	fmt.Println("City:", employee.HomeAddress.City)
	fmt.Println("State:", employee.HomeAddress.State)
	fmt.Println("Postal Code:", employee.HomeAddress.PostalCode)
}
