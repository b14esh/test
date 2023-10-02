package main

import "fmt"

func status(name string) {
	grades := map[string]float64{"Alma": 0, "Rohit": 86.5}
	grade := grades[name]
	if grade < 60 {
		fmt.Printf("%s is failing!\n", name)
	}
	if grade > 60 {
		fmt.Printf("%s grade > 60\n", name)
	}
}
func main() {
	status("Alma") //Ключ которому  присвоено 0
	status("Carl") //Не создавали ключ с таким именем
	status("Rohit")
}
