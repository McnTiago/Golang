// função também é chamada de procedimento ou sub rotina
// é uma parte do código, que pega um dado de entrada e dá um dado de saída

package main

import "fmt"

func media(lista []float64) float64 {
	total := 0.0

	for _, valor := range lista {
		total += valor
	}
	return total / float64(len(lista))
}

func main() {
	lista := []float64{98, 93, 77, 82, 83}
	fmt.Println(media(lista))
}

/*func main(){ // programa que calcula a média de uma sala
	lista := []float64{98,93,77,82,73} // lista de notas

	total := 0.0

	for _, valor:= range lista {
		total += valor
	}
	fmt.Println(total/float64(len(lista)))
}*/
