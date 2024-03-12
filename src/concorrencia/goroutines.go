// É uma função capaz de excutar de modo concorrente com outras funções

package main

import "fmt"

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}

func main() {
	go f(0)
	var escreva string
	fmt.Scanln(&escreva)
}
