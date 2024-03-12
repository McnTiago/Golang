// panic: erro do programador, erro execução de tempo
// recover: interrompe o panic

package main

import "fmt"

func main() {
	defer func() {

		x := recover()
		fmt.Println(x)
	}()
	panic("Panico")
}
