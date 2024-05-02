package main

import "fmt"

//1. Создание вложенных структур

type University struct {
	age       int
	yearBased int
	infoShort string
	infoLong  string
}

type Student struct {
	firstName  string
	lastName   string
	university University
}

func main() {
	//2. Создание экземпляра структур с вложением
	stud := Student{
		firstName: "Fedya",
		lastName:  "Petrov",
		university: University{
			yearBased: 1991,
			infoShort: "cool University",
			infoLong:  "very cool University",
		},
	}

	//3. Получение доступа к вложенным полям структур
	fmt.Println("FirstName:", stud.firstName)
	fmt.Println("LastName:", stud.lastName)
	fmt.Println("Year based Uni:", stud.university.yearBased)
	fmt.Println("Long info:", stud.university.infoLong)

}
