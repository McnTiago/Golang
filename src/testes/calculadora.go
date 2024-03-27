package main

import "fmt"

func main() {
	sm := Soma(1, 2, 3)
	m := Multiplica(10, 10)
	sb := Subtrai(2, 5)
	d := Divide(5, 50)
	fmt.Print(sm, m, sb, d)
}

func Soma(i ...int) int {
	total := 0
	for _, v := range i {
		total += v
	}
	return total
}

func Multiplica(i ...int) int {
	total := 1
	for _, v := range i {
		total = total * v
	}
	return total
}

func Subtrai(i ...int) int {
	total := 0
	for _, v := range i {
		total = v - total
	}
	return total
}

func Divide(i ...int) int {
	total := 1
	for _, v := range i {
		total = v / total
	}
	return total
}
