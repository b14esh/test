package main

import "fmt"

//1. Обьявляем интерфейс декларирующий набор методов под реализацию
type Worker interface {
	Work()
}

//2. Обьявляем структуру, Данная структура это притиндент на удовлетварениея интерфеса Worker
type Employee struct {
	name string
	age  int
}

//3. Реализуем метод work(), что бы структура Employee удовлетворяла интерфейсу Worker

func (e Employee) Work() {
	fmt.Println("Now Employee with name", e.name, "is working!")
}

//4. Адавайте создадим фукциию которая будет принимать аргументы типа Worker и что то с ними делать
func Describer(w Worker) {
	fmt.Printf("Interface with type %T and value %v", w, w)
}

//6.Обьявляем структуру, Данная структура это притиндент на удовлетварениея интерфеса Worker
type Student struct {
	name         string
	courseNumber int
}

func (s Student) WantToEat() {
	fmt.Println("Student with name", s.name, "want to eat")
}

func (s Student) Work() {
	fmt.Println("Now Employee with name", s.name, "is working!")
}

func main() {
	//5. Создадим экземпляр Employee
	emp := Employee{"Bob", 34}
	var workerEmployee Worker = emp //присваиваем сотрудника в переменную типа Worker
	workerEmployee.Work()
	Describer(workerEmployee) // В результате видим что тип интерфеса определяется структурой его реализующей
	//а значение интерфеса - это сооветсвенно значение состояния структуры

	//7. Создадим экземпляр Student
	std := Student{"Mike", 19}
	var workerStudent Worker = std
	workerStudent.Work()
	Describer(workerStudent) //Интерфес принял внутренний тип студент, а значения равно значению экзепляра

	//8. Создадим набор тех, кто умеет Work()
	var workers []Worker = []Worker{workerStudent, workerEmployee}
	for _, worker := range workers {
		Describer(worker) // Данная функция вызывается у разных экземпляров благодря тому, кто для ее вызова
		//экземпляру нужно удовлетворить некому контракту (поведенческому паттерну). Если структура экземпляра поддерживает
		// данный паттерн - то у экземпляра 100% можно вызвать все необходимые для этого методы.
	}

}
