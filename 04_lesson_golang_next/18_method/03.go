package main

import "fmt"

type University struct {
	city string
	name string
}

// Пример наследования
//1. Данный метод явно связон с структурой University
func (u *University) FullinfoUniversity() {
	fmt.Printf("Uni name: %s and City: %s", u.name, u.city)
}

//2. В структуру Proffesor встроены поля структуры University (переходят и все методы тоже)
type Professor struct {
	fullname string
	age      int
	University
}

//1. Анонимные методы
func main() {

	p := Professor{
		fullname: "Boris Bobrov",
		age:      150,
		University: University{
			city: "Moskow",
			name: "BMSTU",
		},
	}
	//3. Вызываем метод структуры University через экземпляр профессора
	p.FullinfoUniversity()
}
