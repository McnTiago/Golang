//Método é uma função associada a um tipo particular
//Isto é em POO, objeto é um valor (variável) e p método
//É uma função associada a esse objeto

package main

import "fmt"

type retangulo struct {
	comprimento, altura int
}

//Esse método 'area' possui um tipo 'retangulo'
func (r retangulo) area() int {
	return r.comprimento * r.altura
}

// Métodos podem ser definidos por qualquer tipo de receptor
// ónteiro ou valor. Aqui temos um exemplo de valor.

func (r retangulo) perimetro() int {
	return 2*r.comprimento + 2*r.altura
}

func main() {
	r := retangulo{comprimento: 50, altura: 25}

	fmt.Println("Area: ", r.area())
	fmt.Println("perimetro: ", r.perimetro())
}
