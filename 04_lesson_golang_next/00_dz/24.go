//Хитрый валидатор паролей LOLKEK9000
//В наше время вопрос обеспечения безопасности в сети стоит достаточно остро, поэтому вам поручено разработать сверх-интеллектуальную систему проверки паролей с замысловатым названием LOLKEK9000.

//Пользователь работает с вашей системой следующим образом. Он вводит 2 строки: пароль и еще раз тот же самый пароль.

//В этот момент в игру вступает LOLKEK9000 и если пароли не удовлетворяют запросам валидатора пользователю придется ввести пароли еще раз (и так будет продолжаться до тех пор, пока не будут выполнены все требования сверхинтеллектуальной системы).

//LOLKEK9000 умеет делать следующее:

//если первый пароль из пары, которую ввел пользователь, короче 8 символов , программа выводит "Слишком короткий пароль!" и заново считывает пару паролей.
//если же первый пароль больше (или равен) 8 символов , но в нем содержатся подстроки "123" или "qwe" - программа выводит "Слишком простой пароль!" и заново считывает пару паролей.
//если же обе вышеуказанные проверки успешно пройдены, но пароли между собой не совпадают - программа выводит "Введенные пароли различаются!" и заново считывает пару паролей.
//если все пункты выше пройдены успешно, программа выводит "Ну наконец-то!" и завершается.

//Ввод:          Вывод:
//star_wars9000 Ну наконец-то!
//star_wars9000

//Ввод:       Вывод:
//derek123    Слишком простой пароль!
//derek123    Ну наконец-то!
//derek9988
//derek9988

package main

import (
	"fmt"
	"strings"
)

func main() {

	var password string
	var save_input string
outline:
	for {
		fmt.Scanln(&password)
		fmt.Scanln(&save_input)

		//if password == "" {
		//	fmt.Println("Вы ничего не ввели!")
		//}

		if len(password) < 8 {
			fmt.Println("Слишком короткий пароль!")
		}

		if strings.Contains(password, save_input) == false && len(password) > 8 {
			fmt.Println("Введенные пароли различаются!")
		}

		if len(password) >= 8 && strings.Contains(password, "qwe") == true || len(password) >= 8 && strings.Contains(password, "123") == true {
			fmt.Println("Слишком простой пароль!")

		}

		if strings.Contains(password, save_input) == true && strings.Contains(password, "qwe") == false && strings.Contains(password, "123") == false && len(password) > 8 {
			fmt.Println("Ну наконец-то!")
			break outline
		}
	}

}
