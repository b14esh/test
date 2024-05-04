package main

import "fmt"

type Employee struct {
	name     string
	position string
	salary   int
	carrency string
}

//1. Методы - функция, привязанные к определенным структурам
// func (s truct) MethodName(parameters type) {}
// Reciever - получатель метода

func (e Employee) Displayinfo() {
	fmt.Println("Name:", e.name)
	fmt.Println("Position:", e.position)
	fmt.Printf("SalaryL: %d %s\n", e.salary, e.carrency)
}

func main() {

	emp := Employee{
		name:     "Bob",
		position: "Senior golang developer",
		salary:   3000,
		carrency: "USD",
	}
	//2. Вызов метода
	emp.Displayinfo()

}
