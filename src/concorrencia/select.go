//Select funciona como um Switch para canais.
// Select permite que você aguarde operações de vários canais.
// Combinar Goroutines e canais no select é um recurso poderoso do Go.

package main

// Cada canal recebera um valor após algum tempo, para simular, por exemplo,
// o bloqueio de operações RPC em execução em Goroutines concorrentes
import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() { //Recebemos os valores "um" e depois "dois" conforme esperado.
		time.Sleep(1 * time.Second)
		c1 <- "um"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "dois"
	}()

	for i := 0; i < 2; i++ {
		select { // Usaremos select para aguardar esses
		// dois valores simultaneamente, imprimento cada um à medida que chegam.
		case msg1 := <-c1:
			fmt.Println("receba", msg1)
		case msg2 := <-c2:
			fmt.Println("receba", msg2)
		}
	}
}
