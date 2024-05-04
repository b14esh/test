package main

import "fmt"

//0. Структура - это явно декларированый зименованный набор СОСТОЯНИЙ
//Структура, исходя из своего описания, отвечает на вопрос - из ЧЕГО я должен состоять,
// чтобы считаться тем ТИПОМ, который деларируется структурой?
// Структура - описывает ПАТТЕРН СОСТОЯНИЯ.

//1. Интерфейс - это что ?
// Это явно декларированный набор сигнатур ПОВЕДЕНИЕЙ (чаще всего об виде набора методов), удовлетворив который,
// можно считатся типом, который декларирует интерфес.
// Интерфейс в GO декларирует ПОЛУ-ОБСТРАКТНЫЙ КЛАСС.
// Интерфес описывает патерн поведения.
// Отвечает на вопрос что я ДОЛЖЕН УМЕТЬ, что бы считатся тем ТИПОМ, который декламирует интерфес

//2. Обьявление интерфейса
type FigureBuilder interface {
	//Набор сигнатур методов, которые необходимо реализовать в структуре-претендента
	//Во-первых, у претендента должен быть метод Area() возвращающий int
	//Во-вторых, у претендента должен быть метод Perimeter() возвращающий int
	Area() int
	Perimeter() int
}

//3. декларируем прентедентов
//3.1 Первый прентендент - это прямоугольник
// У него есть два метода
//Area() int
//Perimeter() int
//когда два этих метода реализованы говорят RECTANGLE УДОВЛЕТВОРЯЕТ УСЛОВИЯМ ИНТЕРФЕСА FigureBuilder
// RECTANGLE РЕАЛИЗУЕТ ИНТЕРФЕЙС FigureBuilder
type Rectangle struct {
	length, width int
}

func (r Rectangle) Area() int {
	return r.length * r.width
}

func (r Rectangle) Perimeter() int {
	return 2 * (r.length + r.width)
}

//3.2 Второй претендент это окружнсот
// У него есть два метода
//Area() int
//Perimeter() int
//когда два этих метода реализованы говорят Circle УДОВЛЕТВОРЯЕТ УСЛОВИЯМ ИНТЕРФЕСА FigureBuilder
// Circle РЕАЛИЗУЕТ ИНТЕРФЕЙС FigureBuilder
type Circle struct {
	radius int
}

func (c Circle) Area() int {
	//Pi округили до 3
	return 3 * c.radius * c.radius
}

func (c Circle) Perimeter() int {
	return 2 * 3 * c.radius
}

func main() {
	//4. Создадим два экземпляра
	r1 := Rectangle{10, 20}
	r2 := Rectangle{30, 50}
	r3 := Rectangle{1, 1}
	c1 := Circle{5}
	c2 := Circle{10}
	c3 := Circle{15}

	//5. Посчитаем общую прлощадь этих фигур
	//total := r.Area() + c.Area()
	rectangle := []Rectangle{r1, r2, r3}
	totalSumAreaOfrectangle := 0
	for _, rect := range rectangle {
		totalSumAreaOfrectangle += rect.Area()
	}

	circle := []Circle{c1, c2, c3}
	totalSumAreaOfCircle := 0
	for _, rect := range circle {
		totalSumAreaOfCircle += rect.Area()
	}
	fmt.Println("Total area is:", totalSumAreaOfrectangle+totalSumAreaOfCircle)

	//6. Теперь воспользуемся ипнтерфесом
	//figures := make([]FigureBuilder, 0, 10) // обьявляю слайс удовлетворяющий интерфейсу FigureBuilder
	// С другой стороны, кажется чот это слайс- каких то  пределенных типов

	//7. Получаем общую площадь
	figures := []FigureBuilder{r1, r2, r3, c1, c2, c3} //обьявим явно
	total := 0
	for _, fig := range figures {
		total += fig.Area()
	}
	fmt.Println("Total area is:", total)
	//8. Пояснение - так как каждый экземпляр удовлетворяет интерфейсу FigureBuilder (Обьявляющий полу-абстрактный тип)
	// у каждого из слайсов можно вызвать 100% Aria вернет int или Perimeter вернет int
}
