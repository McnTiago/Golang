package main

import "fmt"

const tempK = 373.0

func main() {

	tempC := (tempK - 273)

	fmt.Printf("O ponto de ebulição da agua em ºK é %g e o ponto em ºC é %g", tempK, tempC)

}
