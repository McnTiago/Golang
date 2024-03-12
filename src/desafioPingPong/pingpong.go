package main

import (
	"fmt"
	"time"
)

func ping(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
		time.Sleep(1 * time.Second)
	}
}
func pong(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
		time.Sleep(1 * time.Second)
	}
}

func imprimir(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {

	var c chan string = make(chan string)

	go ping(c)
	go pong(c)
	go imprimir(c)

	var entrada string
	fmt.Scanln(&entrada)
}
