package main

import "fmt"

// Utilizando for
/*func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}*/

// Utilizando If
/*func main() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, " é Par")
		} else {
			fmt.Println(i, " é Impar")
		}
	}
}*/

// Utilizando IfElse
/*func main() {
	for i := 1; i <= 5; i++ {
		if i == 0 {
			fmt.Println("Zero")
		} else if i == 1 {
			fmt.Println("Um")
		} else if i == 2 {
			fmt.Println("Dois")
		} else if i == 3 {
			fmt.Println("Três")
		} else if i == 4 {
			fmt.Println("Quatro")
		} else if i == 5 {
			fmt.Println("Cinco")
		}
	}
}*/

// Utilizando Swith
/*func main() {
	for i := 0; i <= 10; i++ {
		switch i {
		case 0:
			fmt.Println("Zero")
		case 1:
			fmt.Println("Um")
		case 2:
			fmt.Println("Dois")
		case 3:
			fmt.Println("Três")
		case 4:
			fmt.Println("Quatro")
		case 5:
			fmt.Println("Cinco")
		}
	}
}*/

// Utilizando for para funcionar como um while
/*func main() {

	i := 0

	for i < 20 {
		fmt.Println(i, "É menor que 20.")
		i++
	}
}*/

// Utilizando o Continue (Ele pula a linha)
func main() {

	for x := 0; x < 20; x++ {
		if x == 3 {
			continue
		}
		fmt.Println(x)
	}
}
