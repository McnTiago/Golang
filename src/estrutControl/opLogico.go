package main

import "fmt"

//Utilizando OU ||
/*func main() {
	x := 4
	if x == 2 || x == 3 {
		fmt.Println("Sim, x é 2 ou 3")
	} else {
		fmt.Println("X não é nem 2 nem 3")
	}
}*/

// Utilizando %%
func main() {
	x := 12
	if x%2 == 0 && x%3 == 0 {
		fmt.Println("X é multiplo de 2 e 3")
	} else {
		fmt.Println("X não é multiplo de 2 e 3")
	}
}
