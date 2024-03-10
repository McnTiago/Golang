// Interfaces são coleções de métodos

package main

import (
	"fmt"
	"math" // Utilizar o pi
)

// Aqui temos uma interface usada para formas geométricas.
type geometria interface {
	area() float64
	perimetro() float64
}

// Para nossa aula vamos usar essa interface nos
// tipos QUADRADO e CIRCULO
// área de um quadrado lado2 ou lado*lado
// perímetro = soma dos lados
type quadrado struct {
	lado float64
}
type circulo struct { //área círculo = (pi)*raio2 érímetro do círculo 2*r*(pi)
	raio float64
}

// Para usar uma interface em Go, só precisamos
// usar todos os métodos da interface. Aqui nós
// usaremos 'geometria' em 'quadrado's.
func (q quadrado) area() float64 {
	return q.lado * q.lado
}
func (q quadrado) perimetro() float64 {
	return 4 * q.lado
}

// A implementação dos 'círculos's.
func (c circulo) area() float64 {
	return math.Pi * c.raio * c.raio
}
func (c circulo) perimetro() float64 {
	return 2 * math.Pi * c.raio
}

// Se a variavel tem um tipo interface, então nós podemos chamar
// métodos que estão na interface nomeada. Aqui temos uma
// função geométrica 'medida' tomando vantagem desse
// trabalho em qualquer 'geometria'
func medir(g geometria) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimetro())
}

func main() {
	q := quadrado{lado: 25}
	c := circulo{raio: 100}

	// Os tipos de estrutura 'circulo' e 'quadrado' ambos
	// implementam a interface 'geometria' então nós podemos usar
	// instâncias dessas estruturas como argumentos para 'medir'

	medir(q)
	medir(c)

}
