package main

import "fmt"

func main() {

	var (
		name string
		str1 string
		str2 string
		str3 string
		str4 string
	)

	fmt.Println("Введите четыре слова. \nПример: \nМир Привет Земля Луна \nПитон Это Очень Круто")
	fmt.Scan(&str1, &str2, &str3, &str4)
	fmt.Println("Введите ваше имя. \n Пример: \nЖеня\nBestPracticeSchool")
	fmt.Scan(&name)
	fmt.Println("\nВы накалялкали:")
	fmt.Println(str4, "-", name)
	fmt.Println(str3, "-", name)
	fmt.Println(str2, "-", name)
	fmt.Println(str1, "-", name)
}
