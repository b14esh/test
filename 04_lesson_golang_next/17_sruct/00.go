package main

import "fmt"

//1. Структура - заименованный набор полей (состояний), определяющий новый тип данных.
//2. Определение структуры (явное определение)
type Student struct {
	firstName string
	lastName  string
	age       int
}

//3.Если имеется состояние одного типа, можно сделать так:
type AnotherStudent struct {
	firstName, lastName, groupName string
	age, cuorseNumber              int
}

func PrintStudent(std Student) {
	fmt.Println("==============================")
	fmt.Println("FirstName:", std.firstName)
	fmt.Println("LasttName:", std.lastName)
	fmt.Println("Age:", std.age)
	fmt.Println("==============================")

}

//11. Структура с Анонимными полями
type Human struct {
	firstName string
	lastName  string
	string
	int
	bool
}

func main() {
	//4. Создание представителя структуры
	stud1 := Student{
		firstName: "Fedya",
		lastName:  "Petrov",
		age:       21, // Запетую можно убрать но придется перенести скобачку "}", а иначе го туда влепи ; и будет ошибка
	}

	fmt.Println("Student 1:", stud1)

	PrintStudent(stud1)

	//5.Короткае определение структуры:
	stud2 := Student{"Petya", "Ivanov", 19}
	PrintStudent(stud2)

	//6. Что если не указать все поля структуры\
	stud3 := Student{
		firstName: "Vasia",
	}
	PrintStudent(stud3)

	//7. Нулевые значения
	stud4 := Student{}
	PrintStudent(stud4)

	//8. Анонимные структуры (у них нет имени) //Очень часто используется в тестах
	anonStudent := struct {
		age           int
		groupID       int
		proffesorName string
	}{
		age:           23,
		groupID:       2,
		proffesorName: "Allex",
	}

	fmt.Println("ANNON STUDENT", anonStudent)

	//9. Досуп к состояниям и их модификациям
	studVova := Student{"Vova", "Ivanov", 10}
	fmt.Println("Firstname:", studVova.firstName)
	fmt.Println("Lastname:", studVova.lastName)
	fmt.Println("Age:", studVova.age)
	studVova.age += 2 // прибавим 2
	fmt.Println("Age:", studVova.age)

	//10. Иницилизация пустой структуры
	emptyStudent := Student{} //способ раз
	var emptyStudent2 Student // способ два
	PrintStudent(emptyStudent)
	PrintStudent(emptyStudent2)

	//11. Указатели на экземпляры структур
	studPointer := &Student{
		firstName: "Igor",
		lastName:  "Sidorov",
		age:       22,
	}

	fmt.Println("value studPointer", studPointer)
	secondPointer := studPointer
	(*&secondPointer).age += 20
	fmt.Println("valuePointerModify:", studPointer)
	studPointerNew := new(Student)
	fmt.Println(studPointerNew)

	//12. Работа с доступом полям сруктуры
	fmt.Println("Age via(*).age:", (*studPointer).age)
	fmt.Println("Age via(*).age:", studPointer.age) // Отработает правильно, в этот момент не явно происходит разыменование указателя studpinter и запрос соотв поля

	//13. Создание структруы с анонимными полями структуры
	human := &Human{
		firstName: "Bob",
		lastName:  "Jonson",
		string:    "Addition info",
		int:       -1,
		bool:      true,
	}

	fmt.Println(human)
	fmt.Println(human.lastName)
	fmt.Println(human.int)
	fmt.Println(human.string)
}
