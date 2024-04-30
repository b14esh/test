//Голодные игры

//Цверг`хакагны - сверхопасные существа. Они нападают на все введенные слова и откусывают от них по 2 символа.
//В первый раз они откусывают 2 символа в начале слова (первые 2 символа) , потом 2 символа в конце (последние 2 символа).
//И так до тех пор, пока не останется один или два символа.
//Напишите программу, которая будет демонстрировать - как цверг`хакагны нападают на слова.

/*
берегись

берегись
регись
реги
ги


нападают

нападают
падают
пада
да

*/

package main

import "fmt"

func main() {

	//for test

	result := "нападают"
	box := []rune(result)

	/*
		for i := 0; i < len(box); i++ {
			fmt.Printf("%c", box[i])
		}
		fmt.Println()

		for i := len(box) - 1; i >= 0; i-- {
			fmt.Printf("%c", box[i])
		}

		fmt.Println()

		for i := 0; i < len(box); i++ {
			for i := 0; i < len(box); i++ {
				fmt.Printf("%c", box[i])
			}
			fmt.Println()
		}

		for i := 0; i < len(box); i++ {
			for a := 0; a < len(box); a++ {
				for i := 0 + a; i < len(box); i++ {
					fmt.Printf("%c", box[i])
				}
				fmt.Println()

			}
			//fmt.Println()
		}
		fmt.Println()
		fmt.Println()

		for a := 0; a < len(box); a++ {
			for i := 0 + a; i < len(box); i++ {
				fmt.Printf("%c", box[i])
			}
			fmt.Println()

		}

		fmt.Println()
		fmt.Println()

		for x, y := 0, 1; x <= 10 && y <= 12; x, y = x+1, y+2 {
			fmt.Printf("%d + %d = %d\n", x, y, x+y)

		}

		fmt.Println()
		fmt.Println()

		for a := 0; a < len(box); a++ {
			for i, b := 0+a, len(box)-1-a; i < len(box) && b >= 0; i, b = i+1, b-1 {

				fmt.Printf("%c %c", box[i], box[b])
			}
			fmt.Println()

		}
		fmt.Println()
		fmt.Println()

		for a := 0; a < len(box); a++ {
			for i, b := 0+a, len(box)-1-a; i < len(box) && b >= 0; i, b = i+1, b-1 {

				fmt.Printf("%c %c", box[i], box[b])
			}
			fmt.Println()

		}
		fmt.Println()
		fmt.Println()

		for a := 0; a < len(box); a++ {
			for i, b := 0+a, len(box)-1-a; i < len(box) && b >= 0; i, b = i+1, b-1 {

				fmt.Printf("%c %c", box[i], box[b])
			}
			fmt.Println()

		}

		fmt.Println()
		fmt.Println()

		//var x int


	*/
	/*
		for a := 0; a < len(box); a++ {
			for i, b := 0+a, len(box)-1; i < len(box) && b > 0; i, b = i+1, b-a {

				fmt.Printf("%c", box[i])

			}

			fmt.Println()

		}
	*/

	/*
		for a, b := 0, 0; a < len(box) && b < len(box); a, b = a+1, b+1 {
			for i, i2 := a, b; i < len(box) && i2 < len(box); i, i2 = i+1, i2+1 {

				fmt.Printf("%c", box[i])
				//fmt.Printf("%c", box[b])

			}

			fmt.Println()

		}

	*/

	/* уже близко :)

	for a := 0; a < len(box); a++ {
		for i, b := 0+a, len(box)-1; i < len(box) && b > 0; i, b = i+1, b-a {

			if i == len(box) {
				continue
			}

			fmt.Printf("%c", box[i])


		}

		fmt.Println()

	}
	*/

	superbox := []rune{}

	for a := 0; a < len(box); a++ {

		for i, b := 0+a, len(box)-1; i < len(box) && b > 0; i, b = i+1, b-a {

			if i == len(box) {
				continue
			}
			superbox = append(superbox, box[i])
			fmt.Printf("%c", box[i])

		}

		fmt.Printf("%c", box)
		fmt.Println()
	}

	/*

		word := "нападают"
		runes := []rune(word)
		for i := 0; i < len(runes)-1; i += 2 {
			fmt.Println(string(runes[i:]))
		}

	*/

	word := "нападают"
	runes := []rune(word)

	for i := 0; i < len(runes)/2; i += 2 {
		fmt.Println(string(runes[i : len(runes)-i]))
	}

	/*
			word := "нападают"

			runes := []rune(word)

			for i := 0; i < len(runes); i++ {
			 fmt.Println(string(runes[i:]))
			 i++
			}
		   }
	*/

	for i := 0; i < len(runes); i++ {
		fmt.Println(string(runes[i:]))
		i++
	}

	/*
		for i := 0; i < len(box); i++ {
			fmt.Printf("%c", box[i])
			fmt.Println()
		}

	*/

	/*

		/*
		for i := 0; i < len(box)-1 && y > i; i++ {
			if i%2 == 1 {
				y--
			}
			fmt.Println(string(box[i:y]))

		}
	*/

	/*
		first := box[0]
		second_first := box[1]
		end := box[len(box)-1]
		second_end := box[len(box)-2]
		y := len(box)

		fmt.Println(string(first), string(second_first), string(end), string(second_end))

		for i := 0; i < len(box) && y > i; i++ {
			fmt.Println(string(box[i:y]))
			i++
			y--

		}

	*/

	/*
		input := bufio.NewScanner(os.Stdin)

		for {
			if input.Scan() {
				result := input.Text()

				if result == "" {
					break
				}

				box := []rune(result)

				fmt.Println()

				for i := 0; i < len(box); i++ {
					for i := 0; i < len(box); i++ {
						fmt.Printf("%c", box[i])
					}
					fmt.Println()
				}

				break
			}

		}
	*/

}
